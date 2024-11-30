package dating

import (
	"github.com/uchupx/dating-api/config"
	"github.com/uchupx/dating-api/pkg/jwt"
	"github.com/uchupx/dating-api/src/service"
)

func (i *Dating) AuthService(conf *config.Config) *service.AuthService {
	if i.authService == nil {
		i.authService = &service.AuthService{
			UserRepo:         i.UserRepo(conf),
			JWT:              i.JWTService(conf),
			Redis:            i.RedisClient(conf),
			ClientRepo:       i.ClientRepo(conf),
			RefreshTokenRepo: i.RefreshTokenRepo(conf),
		}
	}

	return i.authService
}

func (i *Dating) UserService(conf *config.Config) *service.UserService {
	if i.userService == nil {
		i.userService = &service.UserService{
			UserRepo: i.UserRepo(conf),
		}
	}

	return i.userService
}

func (i *Dating) JWTService(conf *config.Config) jwt.CryptService {
	if i.jwtService == nil {
		rsa := jwt.NewCryptService(jwt.Params{
			PrivateKey: conf.RSA.Private,
			PublicKey:  conf.RSA.Public,
		})

		i.jwtService = rsa
	}

	return i.jwtService
}