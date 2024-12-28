package config

import (
	"os"

	"github.com/joho/godotenv"
)

type LoadedEnv struct {
	Environment string
	DbUrl       string
}

func LoadConfig() LoadedEnv {
	if os.Getenv("ENVIRONMENT") == "" {
		err := godotenv.Load(".env")

		if err != nil {
			panic("Error loading .env file")
		}
	}

	return LoadedEnv{
		Environment: os.Getenv("ENVIRONMENT"),
		DbUrl:       os.Getenv("DB_URL"),
	}
}
