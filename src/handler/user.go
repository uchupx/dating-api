package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/uchupx/dating-api/pkg/errors"
	"github.com/uchupx/dating-api/src/dto"
	"github.com/uchupx/dating-api/src/middleware"
	"github.com/uchupx/dating-api/src/service"
)

type UserHandler struct {
	UserService *service.UserService
}

func (h *UserHandler) InitRoutes(e *echo.Echo, m *middleware.Middleware) {
	e.GET("/user/me", h.Me, m.Authorization)
	e.GET("/user/random", h.GetRandom, m.Authorization)
	e.POST("/user/reaction", h.Reaction, m.Authorization)
}

func (h *UserHandler) Me(c echo.Context) error {
	userID := c.Request().Context().Value("userData").(string)

	res, err := h.UserService.FindUserByID(c.Request().Context(), userID)
	if err != nil {
		return responseError(c, err)
	}

	return c.JSON(200, res)
}

func (h *UserHandler) Update(c echo.Context) error {
	return nil
}

func (h *UserHandler) Get(c echo.Context) error {
	return nil
}

func (h *UserHandler) GetRandom(c echo.Context) error {
	res, err := h.UserService.FindRandomUser(c.Request().Context())
	if err != nil {
		return responseError(c, err)
	}

	return c.JSON(200, res)
}

func (h *UserHandler) Reaction(c echo.Context) error {
	var req dto.ReactionRequest
	if err := c.Bind(&req); err != nil {
		return responseError(c, &errors.ErrorMeta{HTTPCode: 400, Message: "invalid request"})
	}

	res, err := h.UserService.Reaction(c.Request().Context(), req)
	if err != nil {
		return responseError(c, err)
	}

	return c.JSON(200, res)
}
