package main

import (
	"target/cmd"
	"target/config"
	"target/db"
	"target/internal/handler"
	"target/internal/repository"
	"target/internal/service"
)

func main() {
	loadedConfig := config.LoadConfig()
	connDb := db.ConnDB(db.Config{
		Environment: loadedConfig.Environment,
		Driver: loadedConfig.DBDriver,
		Host: loadedConfig.DBHost,
		Port: loadedConfig.DBPort,
		User: loadedConfig.DBUser,
		Password: loadedConfig.DBPass,
		Database: loadedConfig.DBName,
		SSLMode: loadedConfig.DBSSLMode,
	})

	goalsRepository := repository.NewGoalsRepository(connDb)

	goalsService := service.NewGoalsService(*goalsRepository)

	goalsHandler := handler.NewGoalHandler(*goalsService)

	cmd.Server(goalsHandler)
}
