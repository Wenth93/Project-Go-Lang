package main

import (

	// Framework V4
	"fmt"
	"project/config"
	"project/handlers"
	"project/repos"
	"project/services"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()

	// load config
	config := config.Load()
	userRepo := repos.NewUserRepository(config.DbConn)
	userService := services.NewUserService(userRepo)

	// API status
	healthHandler := handlers.NewHealthHandler()
	server.GET("/live", healthHandler.IsAlive)

	// User handler
	userHandler := handlers.NewUserHandler(userService)

	// Register POST /users endpoint
	server.POST("/users", userHandler.Post)

	// Login POST /users endpoint
	server.POST("/login", userHandler.Login)

	if err := server.Start(":1323"); err != nil {
		fmt.Println(err)
	}
}
