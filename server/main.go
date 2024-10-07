package main

import (
	"target/cmd"
	"target/config"
	"target/db"
	"target/internal/module"
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

	goalsModule := module.NewGoalsModule(connDb)

	cmd.Server(&goalsModule.GoalsHandler)
}
