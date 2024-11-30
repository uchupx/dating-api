package dating

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/uchupx/dating-api/config"
	"github.com/uchupx/dating-api/src/handler"
	"github.com/uchupx/dating-api/src/middleware"
)

type Dating struct {
	middleware *middleware.Middleware
}

func (i *Dating) isDebug(conf *config.Config) bool {
	return conf.App.Env == "development"
}

func (i *Dating) Middlware(conf *config.Config) {
	// i.middlware = &middleware.Middleware{}
}

func (i *Dating) InitRoutes(conf *config.Config, e *echo.Echo) {
	routes := []handler.BaseHandler{
		&handler.Handler{},
	}
	i.Middlware(conf)
	// e.Use(i.middlware.Logger)
	// eMiddleware.CORSWithConfig(eMiddleware.CORSConfig{})
	e.Use(i.middleware.Recover)

	for _, route := range routes {
		route.InitRoutes(e, i.middleware)
	}

	if i.isDebug(conf) {
		for _, route := range e.Routes() {
			fmt.Printf("Route | Method: %s | Path: %s | Caller: %s \n", route.Method, route.Path, route.Name)
		}
	}
}
