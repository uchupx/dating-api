package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/uchupx/dating-api/src/dto"
	"github.com/uchupx/dating-api/src/middleware"
	"github.com/uchupx/dating-api/src/service"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func (a *AuthHandler) InitRoutes(e *echo.Echo, m *middleware.Middleware) {
	e.POST("/token", a.Auth)
	e.POST("/sign-up", a.SignUp)
	e.POST("/client", a.ClientAdd)
	e.POST("/refresh", a.RefreshToken)
	e.GET("/status", func(c echo.Context) error {
		return c.JSON(200, "Is Auth")
	}, m.Authorization)
}

func (a *AuthHandler) Auth(c echo.Context) error {
	var req dto.AuthRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	res, err := a.AuthService.Login(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(200, res)
}

func (a *AuthHandler) SignUp(c echo.Context) error {
	var req dto.SignUpRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	res, err := a.AuthService.SignUp(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(201, res)
}

func (a *AuthHandler) ClientAdd(c echo.Context) error {
	var req dto.ClientPost
	if err := c.Bind(&req); err != nil {
		return err
	}

	res, err := a.AuthService.AddClient(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(201, res)
}

func (a *AuthHandler) RefreshToken(c echo.Context) error {
	var req dto.RefreshTokenRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	res, err := a.AuthService.RefreshToken(c.Request().Context(), req)
	if err != nil {
		return err
	}
	return c.JSON(200, res)
}
