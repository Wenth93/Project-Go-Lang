package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) IsAlive(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Ok"})
}
