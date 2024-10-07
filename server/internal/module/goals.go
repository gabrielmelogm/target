package module

import (
	"database/sql"
	"target/internal/handler"
	"target/internal/repository"
	"target/internal/service"
)

type GoalsModule struct {
	GoalsRepository repository.GoalsRepository
	GoalsService    service.GoalsService
	GoalsHandler    handler.GoalsHandler
}

func NewGoalsModule(conn *sql.DB) *GoalsModule {
	goalsRepository := repository.NewGoalsRepository(conn)

	goalsService := service.NewGoalsService(*goalsRepository)

	goalsHandler := handler.NewGoalHandler(*goalsService)

	return &GoalsModule{
		GoalsRepository: *goalsRepository,
		GoalsService:    *goalsService,
		GoalsHandler:    *goalsHandler,
	}
}
