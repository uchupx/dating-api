package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sirupsen/logrus"
	"github.com/uchupx/kajian-api/pkg/errors"
	"github.com/uchupx/kajian-api/pkg/logger"
)

func (m *Middleware) Recover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		defer func() {
			if r := recover(); r != nil {
				if r == http.ErrAbortHandler {
					panic(r)
				}

				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}

				fileLine := errors.TracePanic()

				meta := errors.ErrorMeta{
					Message: "Internal Server Error",
					Line:    fileLine,
					IsPanic: true,
				}

				logger.Logger.WithFields(logrus.Fields{
					"meta": meta, // Log response body
				}).Error(err)

				c.Error(err)
			}
		}()
		return next(c)
	}
}
