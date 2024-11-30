package middleware

import (
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/uchupx/dating-api/pkg/helper"
)

func (m *Middleware) Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		reg := regexp.MustCompile(`Bearer[\s]`)

		auth = reg.ReplaceAllString(auth, "")
		if strings.TrimSpace(auth) != "" {
			val, err := m.Redis.Get(c.Request().Context(), helper.REDIS_KEY_AUTH)
			if err != nil || val == nil {
				return echo.NewHTTPError(401, "Unauthorized")
			}

			next(c)
		} else {
			return echo.NewHTTPError(401, "Unauthorized")
		}
		return nil
	}
}
