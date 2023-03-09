package handlers

import (
	"context"
	"log"
	"wrapup/database"

	"go.mongodb.org/mongo-driver/bson"
)

type ErrorResponse struct {
    Message string `json:"message"`
}


func CheckUserExists(client *database.Client, userID string) bool {
	// create a MongoDB filter to find a user with the given ID
	filter := bson.M{"_id": userID}
	// create a MongoDB collection to store users
	coll := client.Database("wrapup-base").Collection("users")
	// count the number of users matching the filter
	count, err := coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		return false
	}
	// return true if a user was found, false otherwise
	return count > 0
}
