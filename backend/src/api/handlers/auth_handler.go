package handlers

import (
	"net/http"
	"time"

	"wrapup/database"
	"wrapup/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginRequest struct {
	Username    string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserID primitive.ObjectID `json:"id"`
	Token string `json:"token"`
}

type Claims struct {
	UserID primitive.ObjectID `json:"user_id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type AuthHandler struct {
	services.UserService	
}

func NewAuthHandler(db *database.Client) *AuthHandler {
    return &AuthHandler{UserService: &services.UserServiceImpl{Db: db}}
}

func (au *AuthHandler) Login(c echo.Context) error {
		req := new(LoginRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		}
		user, err := au.AuthenticateUser(req.Username, req.Password)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "invalid email or password"})
		}
		
		// create a JWT token with a 24 hour expiration time
		claims := &Claims{
			UserID: user.ID,
			Role: user.Role,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte("secret_key"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}
		resp := LoginResponse{Token: tokenString, UserID: user.ID}
		return c.JSON(http.StatusOK, resp)
}
