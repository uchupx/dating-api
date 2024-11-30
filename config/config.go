package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var configPath = []string{
	"./",
	"./config/",
}

type Config struct {
	App struct {
		Env  string
		Port string
		Log  string
	}

	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}

	Redis struct {
		Host     string
		Port     string
		Password string
	}

	RSA struct {
		Private string
		Public  string
	}
}

var config *Config

// NewConfig is a constructor for Config
func new() *Config {
	var err error
	for _, c := range configPath {
		err = godotenv.Load(c + ".env")
		if err != nil {
			log.Printf("failed to laod config from path %s.env", c)
			continue
		}
		break
	}
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config = &Config{}

	config.Database.Host = os.Getenv("DATABASE_HOST")
	config.Database.Port = os.Getenv("DATABASE_PORT")
	config.Database.User = os.Getenv("DATABASE_USERNAME")
	config.Database.Password = os.Getenv("DATABASE_PASSWORD")
	config.Database.Name = os.Getenv("DATABASE_NAME")

	config.Redis.Host = os.Getenv("REDIS_HOST")
	config.Redis.Port = os.Getenv("REDIS_PORT")
	config.Redis.Password = os.Getenv("REDIS_PASSWORD")

	config.RSA.Private = os.Getenv("RSA_PRIVATE_KEY")
	config.RSA.Public = os.Getenv("RSA_PUBLIC_KEY")

	config.App.Env = os.Getenv("APP_ENV")
	config.App.Port = os.Getenv("APP_PORT")
	config.App.Log = os.Getenv("APP_LOG")
	return config
}

func GetConfig() *Config {
	if config == nil {
		config = new()
	}

	return config
}
