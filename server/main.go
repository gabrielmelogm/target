package main

import (
	"target/cmd"
	"target/config"
	"target/db"
	"target/db/seeds"
	"target/internal/repository"
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
	seeds := seeds.NewSeedsRun(*goalsRepository)

	seeds.Run()

	cmd.Server()
}
