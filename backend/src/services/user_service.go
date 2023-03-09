package services

import (
	"context"
	"errors"
	"wrapup/database"
	"wrapup/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Define a UserService interface
type UserService interface {
    GetUser(userID primitive.ObjectID) (*models.User, error)
	GetUserByUsernameOrEmail(username string, password string) (*models.User, error)
    UpdateUser(userID primitive.ObjectID, updateDoc bson.M) (*models.User, error)
    CreateUser(user *models.User) error
    DeleteUser(userID primitive.ObjectID) error
    GetAllUsers() ([]*models.User, error)
	AuthenticateUser(username, password string) (*models.User, error) 

}

// Implement UserService using the database.Client
type UserServiceImpl struct {
    Db *database.Client
}

func (us *UserServiceImpl) GetUserByUsernameOrEmail(username, email string) (*models.User, error) {
    ctx := context.Background()
    db := us.Db.Database("wrapup-base")
    usersColl := db.Collection("users")
    var user models.User
    err := usersColl.FindOne(ctx, bson.M{"$or": []bson.M{
        {"username": username},
        {"email": email},
    }}).Decode(&user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (us *UserServiceImpl) GetUser(userID primitive.ObjectID) (*models.User, error) {
    ctx := context.Background()
    collection := us.Db.Database("wrapup-base").Collection("users")
    var user *models.User
    err := collection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (us *UserServiceImpl) UpdateUser(userID primitive.ObjectID, updateDoc bson.M) (*models.User, error) {
    ctx := context.Background()
    db := us.Db.Database("wrapup-base")
    usersColl := db.Collection("users")
    _, err := usersColl.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$set": updateDoc})
    if err != nil {
        return nil, err
    }
    // Fetch updated user
    var updatedUser models.User
    err = usersColl.FindOne(ctx, bson.M{"_id": userID}).Decode(&updatedUser)
    if err != nil {
        return nil, err
    }
    return &updatedUser, nil
}

func (us *UserServiceImpl) CreateUser(user *models.User) error {
	hashedPassword, err := HashPassword(user.Password)
    if err != nil {
        return err
    }
    user.Password = hashedPassword
	user.ID = primitive.NewObjectID()
    ctx := context.Background()
    db := us.Db.Database("wrapup-base")
    _, err = db.Collection("users").InsertOne(ctx, user)
    return err
}

func (us *UserServiceImpl) DeleteUser(userID primitive.ObjectID) error {
    ctx := context.Background()
    _, err := us.Db.Database("wrapup-base").Collection("users").DeleteOne(ctx, bson.M{"_id": userID})
    return err
}

func (us *UserServiceImpl) GetAllUsers() ([]*models.User, error) {
    ctx := context.Background()
    db := us.Db.Database("wrapup-base")
    users, err := db.Collection("users").Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer users.Close(ctx)

    var result []*models.User
    for users.Next(ctx) {
        var user models.User
        if err := users.Decode(&user); err != nil {
            return nil, err
        }
        result = append(result, &user)
    }

    return result, nil
}

func (us *UserServiceImpl) AuthenticateUser(username, password string) (*models.User, error){
	user, err := us.GetUserByUsernameOrEmail(username, username)
	if err != nil {
		return nil, err
	}
	if !CheckPassword(password, user.Password) {
        return nil, errors.New("incorrect password")
   } 
   return user, nil
}
