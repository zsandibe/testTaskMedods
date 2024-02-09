package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   serverConfig
	Database databaseConfig
	Token    tokenConfig
}

type databaseConfig struct {
	MongoUrl       string
	NameDb         string
	NameCollection string
}

type serverConfig struct {
	Host string
	Port string
}
type tokenConfig struct {
	AccessKey       string
	AccessTokenAge  time.Duration
	RefreshTokenAge time.Duration
}

func NewConfig() Config {
	if err := godotenv.Load(".env"); err != nil {
		return Config{}
	}
	durationAccess, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_AGE"))
	if err != nil {
		fmt.Println("Error parsing access token duration", err)
		return Config{}
	}
	durationRefresh, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_AGE"))
	if err != nil {
		fmt.Println("Error converting REFRESH_TOKEN_AGE to integer:", err)
		return Config{}
	}
	return Config{
		serverConfig{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		databaseConfig{
			MongoUrl:       os.Getenv("MONGODB_URL"),
			NameDb:         os.Getenv("NAME_DB"),
			NameCollection: os.Getenv("NAME_COLLECTION"),
		},
		tokenConfig{
			AccessKey:       os.Getenv("ACCESS_KEY"),
			AccessTokenAge:  time.Duration(durationAccess.Seconds()),
			RefreshTokenAge: time.Duration(durationRefresh.Seconds()),
		},
	}
}
