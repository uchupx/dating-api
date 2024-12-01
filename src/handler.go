package dating

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/uchupx/dating-api/config"
	"github.com/uchupx/dating-api/src/handler"
)

type datingHandler struct {
	authHandler *handler.AuthHandler
	userHandler *handler.UserHandler
}

func (i *Dating) AuthHandler(conf *config.Config) *handler.AuthHandler {
	if i.authHandler == nil {
		i.authHandler = &handler.AuthHandler{
			AuthService: i.AuthService(conf),
		}
	}

	return i.authHandler
}

func (i *Dating) UserHandler(conf *config.Config) *handler.UserHandler {
	if i.userHandler == nil {
		i.userHandler = &handler.UserHandler{
			UserService: i.UserService(conf),
		}
	}

	return i.userHandler
}

func (i *Dating) InitRoutes(conf *config.Config, e *echo.Echo) {
	routes := []handler.BaseHandler{
		&handler.Handler{},
		i.AuthHandler(conf),
		i.UserHandler(conf),
	}

	i.InitLogger(conf)
	i.middleware(conf)
	// e.Use(i.middleware.Logger)
	// eMiddleware.CORSWithConfig(eMiddleware.CORSConfig{})
	e.Use(i.apiMiddleware.Recover)

	for _, route := range routes {
		route.InitRoutes(e, i.apiMiddleware)
	}

	if i.isDebug(conf) {
		for _, route := range e.Routes() {
			fmt.Printf("Route | Method: %s | Path: %s | Caller: %s \n", route.Method, route.Path, route.Name)
		}
	}
}
