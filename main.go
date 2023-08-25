package main

import (
	"fmt"

	"github.com/Wenth93/Project-Go-Lang/config"
	"github.com/Wenth93/Project-Go-Lang/handlers"
	"github.com/Wenth93/Project-Go-Lang/repos"
	"github.com/Wenth93/Project-Go-Lang/services"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()

	// load config
	config := config.Load()
	userRepo := repos.NewUserRepository(config.DbConn)
	userService := services.NewUserService(userRepo)

	healthHandler := handlers.NewHealthHandler()
	server.GET("/live", healthHandler.IsAlive)

	// Create user handler
	userHandler := handlers.NewUserHandler(userService)

	// Register POST /users endpoint
	server.POST("/users", userHandler.Post)

	server.POST("/login", userHandler.Login)

	// TODO: Register a new endpoint for POST user

	if err := server.Start(":8080"); err != nil {
		fmt.Println(err)
	}
}
