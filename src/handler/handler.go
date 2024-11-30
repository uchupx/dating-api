package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/uchupx/dating-api/src/middleware"
)

type BaseHandler interface {
	InitRoutes(e *echo.Echo, middleware *middleware.Middleware)
}

type Handler struct {
}

func (h *Handler) InitRoutes(e *echo.Echo, middleware *middleware.Middleware) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})
}
