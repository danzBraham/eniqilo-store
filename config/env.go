package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigEnv struct {
	AppHost string
	AppPort string

	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	DBParams   string

	JWT_SECRET  string
	BCRYPT_SALT int
}

func LoadEnv() *ConfigEnv {
	if err := godotenv.Load(); err != nil {
		log.Panic("error loading .env file")
	}

	return &ConfigEnv{
		AppHost:    os.Getenv("APP_HOST"),
		AppPort:    os.Getenv("APP_PORT"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBParams:   os.Getenv("DB_PARAMS"),
	}
}
