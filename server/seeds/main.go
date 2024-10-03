package main

import (
	"target/config"
	"target/db"
	"target/internal/repository"
	"target/seeds/cmd"
)

func main() {
	loadedConfig := config.LoadConfig()
	connDb := db.ConnDB(db.Config{
		Environment: loadedConfig.Environment,
		Driver:      loadedConfig.DBDriver,
		Host:        loadedConfig.DBHost,
		Port:        loadedConfig.DBPort,
		User:        loadedConfig.DBUser,
		Password:    loadedConfig.DBPass,
		Database:    loadedConfig.DBName,
		SSLMode:     loadedConfig.DBSSLMode,
	})

	goalsRepository := repository.NewGoalsRepository(connDb)
	seedsInstance := cmd.NewSeedsRun(*goalsRepository) 

	seedsInstance.Run()
}
