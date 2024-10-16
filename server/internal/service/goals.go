package service

import (
	"errors"
	db "target/db/sqlc"
	"target/internal/repository"
	"target/internal/request"
	"time"
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

func (g *GoalsService) CreateNewGoalCompletion(goalId string) error {
	now := time.Now()

	offset := (7 - int(now.Weekday())) % 7
	nextSunday := now.AddDate(0, 0, offset)

	offsetFirstDay := -int(now.Weekday())
	firstDayOfWeek := now.AddDate(0, 0, offsetFirstDay)

	goalCompletions, err := g.GoalsRepository.GetGoalCompletionsCountById(goalId, firstDayOfWeek, nextSunday)

	// TODO: Return a simple error
	if goalCompletions.CompletionCount >= int64(goalCompletions.DesiredWeeklyFrequency) {
		return errors.New("Goal already completed this week!")
	}

	err = g.GoalsRepository.CreateNewGoalCompletion(goalId, time.Now())

	return err
}

func (g *GoalsService) GetPendingGoals() ([]repository.GetPendingGoalsResponse, error) {
	now := time.Now()

	offset := (7 - int(now.Weekday())) % 7
	nextSunday := now.AddDate(0, 0, offset)

	offsetFirstDay := -int(now.Weekday())
	firstDayOfWeek := now.AddDate(0, 0, offsetFirstDay)

	goals, err := g.GoalsRepository.GetPendingGoals(firstDayOfWeek, nextSunday)

	return goals, err
}
