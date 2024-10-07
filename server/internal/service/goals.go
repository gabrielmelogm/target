package service

import (
	db "target/db/sqlc"
	"target/internal/repository"
	"target/internal/request"
)

type GoalsService struct {
	GoalsRepository repository.GoalsRepository
}

func NewGoalsService(goalsRepository repository.GoalsRepository) *GoalsService {
	return &GoalsService{
		GoalsRepository: goalsRepository,
	}
}

func (g *GoalsService) CreateNewGoal(createNewGoalDto request.CreateNewGoalRequest) (db.Goal, error) {
	createdGoal, err := g.GoalsRepository.CreateNewGoal(createNewGoalDto)

	return createdGoal, err
}
