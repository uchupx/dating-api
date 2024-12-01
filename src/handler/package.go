package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/uchupx/dating-api/src/middleware"
	"github.com/uchupx/dating-api/src/service"
)

type PackageHandler struct {
	PackageService *service.PackageService
}

func (h *PackageHandler) InitRoutes(e *echo.Echo, m *middleware.Middleware) {
	e.GET("/packages", h.GetPackages, m.Authorization)
	e.POST("/package/:id/purchase", h.Purchase, m.Authorization)
}

func (h *PackageHandler) GetPackages(c echo.Context) error {
	res, err := h.PackageService.GetPackages(c.Request().Context())
	if err != nil {
		return responseError(c, err)
	}
	return c.JSON(200, res)
}

func (h *PackageHandler) Purchase(c echo.Context) error {
	id := c.Param("id")
	res, err := h.PackageService.Purchase(c.Request().Context(), id)
	if err != nil {
		return responseError(c, err)
	}
	return c.JSON(200, res)
}
