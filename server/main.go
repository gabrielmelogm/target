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
		DbUrl:       loadedConfig.DbUrl,
	})

	goalsModule := module.NewGoalsModule(connDb)

	cmd.Server(&goalsModule.GoalsHandler)
}
