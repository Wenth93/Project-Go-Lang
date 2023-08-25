package main

import (
	"fmt"

	"estiam_golang_api_course_finalproject-master/internal/handlers"
	"estiam_golang_api_course_finalproject-master/internal/repos"
	"estiam_golang_api_course_finalproject-master/internal/services"

	"github.com/phramos07/estiam_golang_api_course_finalproject-master/internal/config"

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
