package handlers

import (
	"net/http"

	"github.com/Wenth93/Project-Go-Lang/services"
	"github.com/Wenth93/Project-Go-Lang/types"
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

func (h *UserHandler) Post(ctx echo.Context) error {
	var newUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.Bind(&newUser); err != nil {
		return err
	}

	user := &types.User{
		Username: newUser.Username,
		Password: newUser.Password,
	}

	err := h.userService.CreateNewUser(ctx.Request().Context(), user)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, user)
}

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
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{"token": token})
}
