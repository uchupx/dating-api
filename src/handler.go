package dating

import (
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
