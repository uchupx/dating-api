package dating

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/uchupx/dating-api/config"
	"github.com/uchupx/dating-api/pkg/database/redis"
	"github.com/uchupx/dating-api/pkg/jwt"
	"github.com/uchupx/dating-api/src/handler"
	"github.com/uchupx/dating-api/src/middleware"
	"github.com/uchupx/dating-api/src/repo"
	"github.com/uchupx/dating-api/src/service"

	"github.com/uchupx/kajian-api/pkg/db"
	"github.com/uchupx/kajian-api/pkg/logger"
	"github.com/uchupx/kajian-api/pkg/mysql"
)

type Dating struct {
	// adapter
	db          *db.DB
	redisClient *redis.Redis

	// repo
	userRepo         *repo.UserRepo
	clientRepo       *repo.ClientRepo
	refreshTokenRepo *repo.RefreshTokenRepo

	//middleware
	apiMiddleware *middleware.Middleware

	// handler
	authHandler *handler.AuthHandler
	userHandler *handler.UserHandler

	// service
	userService *service.UserService
	authService *service.AuthService
	jwtService  jwt.CryptService
}

func (i *Dating) DB(conf *config.Config) *db.DB {
	if i.db == nil {
		for idx := 1; idx <= 3; idx++ {
			db, err := mysql.NewConnection(mysql.DBPayload{
				Host:     conf.Database.Host,
				Port:     conf.Database.Port,
				Username: conf.Database.User,
				Password: conf.Database.Password,
				Database: conf.Database.Name,
			})
			if err != nil {
				fmt.Println(err)
				time.Sleep(2 * time.Second)
				fmt.Println("Retry connection")

				if idx == 3 {
					panic(err)
				}

				continue
			}

			db.SetDebug(i.isDebug(conf))
			i.db = db
			break
		}
	}
	return i.db
}

func (i *Dating) RedisClient(conf *config.Config) *redis.Redis {
	if i.redisClient == nil {
		mainKey := "dating-api"

		if i.isDebug(conf) {
			mainKey = mainKey + ":" + conf.App.Env
		}

		redisClient, err := redis.Connection(redis.Config{
			Host:     fmt.Sprintf("%s:%s", conf.Redis.Host, conf.Redis.Port),
			Password: conf.Redis.Password,
			Database: 0,
			MainKey:  mainKey,
		})

		if err != nil {
			panic(err)
		}

		i.redisClient = redisClient

	}

	return i.redisClient
}

func (i *Dating) UserRepo(conf *config.Config) *repo.UserRepo {
	if i.userRepo == nil {
		i.userRepo = repo.NewUserRepo(i.DB(conf))
	}

	return i.userRepo
}

func (i *Dating) ClientRepo(conf *config.Config) *repo.ClientRepo {
	if i.clientRepo == nil {
		i.clientRepo = repo.NewClientRepo(i.DB(conf))
	}

	return i.clientRepo
}

func (i *Dating) RefreshTokenRepo(conf *config.Config) *repo.RefreshTokenRepo {
	if i.refreshTokenRepo == nil {
		i.refreshTokenRepo = repo.NewRefreshTokenRepo(i.DB(conf))
	}

	return i.refreshTokenRepo
}

func (i *Dating) middleware(conf *config.Config) *middleware.Middleware {
	if i.apiMiddleware == nil {
		i.apiMiddleware = middleware.New(middleware.Config{
			Redis: i.redisClient,
		})
	}

	return i.apiMiddleware
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

func (i *Dating) InitLogger(conf *config.Config) {
	logConf := logger.LogConfig{
		Path:    conf.App.Log,
		NameApp: "dating-api",
	}

	logConf.SetDebug(i.isDebug(conf))
	logger.InitLog(logConf)
}

func (i *Dating) isDebug(conf *config.Config) bool {
	return conf.App.Env == "development"
}
