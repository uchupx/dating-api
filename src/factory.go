package dating

import (
	"fmt"
	"time"

	"github.com/uchupx/dating-api/config"
	"github.com/uchupx/dating-api/pkg/database/redis"
	"github.com/uchupx/dating-api/src/middleware"

	"github.com/uchupx/kajian-api/pkg/db"
	"github.com/uchupx/kajian-api/pkg/logger"
	"github.com/uchupx/kajian-api/pkg/mysql"
)

type Dating struct {
	// adapter
	db          *db.DB
	redisClient *redis.Redis

	datingRepo    //----repo
	datingHandler //----handler
	datingService //----service

	//middleware
	apiMiddleware *middleware.Middleware
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

func (i *Dating) middleware(conf *config.Config) *middleware.Middleware {
	if i.apiMiddleware == nil {
		i.apiMiddleware = middleware.New(middleware.Config{
			Redis: i.redisClient,
		})
	}

	return i.apiMiddleware
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
