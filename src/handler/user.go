package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/uchupx/dating-api/src/middleware"
	"github.com/uchupx/dating-api/src/service"
)

type UserHandler struct {
	UserService *service.UserService
}

func (h *UserHandler) InitRoutes(e *echo.Echo, m *middleware.Middleware) {
	e.GET("/users/me", h.Me, m.Authorization)
}

func (h *UserHandler) Me(c echo.Context) error {
	userID := c.Request().Context().Value("userData")
	return c.JSON(200, userID)
}
