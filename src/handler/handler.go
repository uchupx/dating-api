package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/uchupx/dating-api/pkg/errors"
	"github.com/uchupx/dating-api/src/dto"
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

func responseError(c echo.Context, err *errors.ErrorMeta) error {
	return c.JSON(err.HTTPCode, dto.Response{
		Status:  err.HTTPCode,
		Message: fmt.Sprintf("error: %s", err.Message),
	})
}
