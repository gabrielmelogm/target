package config

import (
	"os"
	"github.com/joho/godotenv"
) 

type LoadedEnv struct {
	Environment string
	DBDriver    string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPass      string
	DBName      string
	DBSSLMode   string
}

func LoadConfig() LoadedEnv {
	if os.Getenv("ENVIRONMENT") == "" {
		err := godotenv.Load(".env")

		if err != nil {
			panic("Error loading .env file")
		}
	}

	return LoadedEnv{
		Environment:  os.Getenv("ENVIRONMENT"),
		DBDriver:     os.Getenv("DB_DRIVER"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBUser:       os.Getenv("DB_USER"),
		DBPass:       os.Getenv("DB_PASS"),
		DBName:       os.Getenv("DB_NAME"),
		DBSSLMode:    os.Getenv("DB_SSL_MODE"),
	}
}

