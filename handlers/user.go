package handlers

import (
	"net/http"
	"project/services"
	"project/types"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

// Route for Creation New User
func (h *UserHandler) Post(ctx echo.Context) error {
	var newUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.Bind(&newUser); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request data"})
	}

	user := &types.User{
		Username: newUser.Username,
		Password: newUser.Password,
	}

	err := h.userService.CreateNewUser(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	return ctx.JSON(http.StatusCreated, user)
}

// Route for User Login
func (h *UserHandler) Login(ctx echo.Context) error {
	var loginUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.Bind(&loginUser); err != nil {
		return err
	}

	token, err := h.userService.Authenticate(ctx.Request().Context(), loginUser.Username, loginUser.Password)
	if err != nil {
		// Handling login errors
		switch {
		case strings.Contains(err.Error(), "errorUserNotFound"):
			return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Wrong username"})
		case strings.Contains(err.Error(), "errorPassword"):
			return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Wrong password"})
		case strings.Contains(err.Error(), "errorToken"):
			return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Error when creating a session"})
		default:
			return err
		}
	}

	return ctx.JSON(http.StatusOK, map[string]string{"token": token})
}
