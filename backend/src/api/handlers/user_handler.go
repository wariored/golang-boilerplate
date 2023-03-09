package handlers

import (
	"log"
	"net/http"
	"wrapup/database"
	"wrapup/models"
	"wrapup/services"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


type UserRegistrationRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Email    string `json:"email"`
}

type UserUpdateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"` 
	Role string `json:"role"`
}

type UserResponse struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Role string `json:"role"`
}


type UserHandler struct {
	services.UserService
}

func NewUserHandler(db *database.Client) *UserHandler {
    return &UserHandler{UserService: &services.UserServiceImpl{Db: db}}
}


func parseUserID(userID string) primitive.ObjectID {
	parsedId, _ := primitive.ObjectIDFromHex(userID)
	return parsedId
}


func (uh *UserHandler) GetUser(c echo.Context) error {
	userID := parseUserID(c.Param("userID"))
	user, err := uh.UserService.GetUser(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}
	return c.JSON(http.StatusOK, &UserResponse{ID: userID.Hex(), Username: user.Username, Email: user.Email, Role: user.Role})
}

func (uh *UserHandler) UpdateUser(c echo.Context) error {
	userID := c.Param("userID")
	var userReq UserUpdateRequest
	if err := c.Bind(&userReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
	}

	// Create update document
	updateDoc := bson.M{}
	if userReq.Username != "" {
		updateDoc["username"] = userReq.Username
	}
	if userReq.Password != "" {
		updateDoc["password"] = userReq.Password
	}
	if userReq.Email != "" {
		updateDoc["email"] = userReq.Email
	}
	if userReq.Role != "" {
		updateDoc["role"] = userReq.Role
	}

	// Perform update
	updatedUser, err := uh.UserService.UpdateUser(parseUserID(userID), updateDoc)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	userResponse := UserResponse{
		ID:       updatedUser.ID.Hex(),
		Username: updatedUser.Username,
		Email:    updatedUser.Email,
		Role:     updatedUser.Role,
	}
	return c.JSON(http.StatusOK, &userResponse)
}

func (uh *UserHandler) CreateUser(c echo.Context) error {
	var userRequest UserRegistrationRequest
	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	// check if the username or email is already in use
	if _, err := uh.UserService.GetUserByUsernameOrEmail(userRequest.Username, userRequest.Email); err == nil {
		return c.JSON(http.StatusConflict, "username or email already in use")
	}
	user := models.User{
		Username: userRequest.Username,
		Password: userRequest.Password,
		Email: userRequest.Email,
		Role: "user",
	}
	err := uh.UserService.CreateUser(&user)
	log.Println("error: ", err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusCreated, "user created successfully")
}

func (uh *UserHandler) DeleteUser(c echo.Context) error {
	userID := c.Param("userID")

	if err := uh.UserService.DeleteUser(parseUserID(userID)); err != nil {
		return c.JSON(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, "user deleted successfully")
}

func (uh *UserHandler) GetAllUsers(c echo.Context) error {
	log.Println("ok ok ok")
	users, err := uh.UserService.GetAllUsers()
	if err != nil { 
		return c.JSON(http.StatusInternalServerError, "error fetching users from database")
	}

	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, UserResponse{
			ID: user.ID.Hex(),
			Username: user.Username,
			Email: user.Email,
			Role: user.Role,
		})
	}

	return c.JSON(http.StatusOK, userResponses)
}
