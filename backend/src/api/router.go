package api

import (
	"wrapup/api/handlers"
	"wrapup/api/middlewares"
	"wrapup/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Router creates a new HTTP router
func Router(client *database.Client) *echo.Echo {
	e := echo.New()

	userHandler :=  handlers.NewUserHandler(client)
	authHandler := handlers.NewAuthHandler(client)
	//messageHandler := handlers.MessageHandler{Db: client}
		
	// Authentication routes
	e.POST("/login", authHandler.Login)
	e.POST("/users", userHandler.CreateUser)

	// auth routes
	authGroup := e.Group("")
	authGroup.Use(middlewares.AuthMiddleware)
	authGroup.GET("/users/:userID", userHandler.GetUser)
	authGroup.PUT("/users/:userID", userHandler.UpdateUser)

	// admin routes
	adminGroup := authGroup.Group("")
	adminGroup.Use(middlewares.RoleAccessMiddleware("admin"))
	adminGroup.GET("/users", userHandler.GetAllUsers)
	adminGroup.DELETE("/users/:userID", userHandler.DeleteUser)

	
	// websocket route
	//e.GET("/ws/messages", websocket.Handler(messageHandler.SendMessage))

	// add middleware to Echo instance
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}
