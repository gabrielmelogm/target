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
		DbUrl:       loadedConfig.DbUrl,
	})

	goalsRepository := repository.NewGoalsRepository(connDb)
	seedsInstance := cmd.NewSeedsRun(*goalsRepository)

	seedsInstance.Run()
}
