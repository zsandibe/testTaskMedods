package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoUrl       string
	NameDb         string
	NameCollection string
	AccessKey      string
	RefreshKey     string
}

func NewConfig() Config {
	if err := godotenv.Load(filepath.Join(".", ".env")); err != nil {
		fmt.Printf("Error loading .env file: %v", err)
		return Config{}
	}

	return Config{
		MongoUrl:       os.Getenv("MONGODB_URL"),
		NameDb:         os.Getenv("NAME_DB"),
		NameCollection: os.Getenv("NAME_COLLECTION"),
		AccessKey:      os.Getenv("ACCESS_KEY"),
	}
}
