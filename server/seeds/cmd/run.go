package cmd

import (
	"fmt"
	db "target/db/sqlc"
	"target/internal/repository"
	"target/internal/request"
	"time"
)

type SeedsRun struct {
	GoalsRepository repository.GoalsRepository
}

func NewSeedsRun(goalsRepository repository.GoalsRepository) *SeedsRun {
	return &SeedsRun{
		GoalsRepository: goalsRepository,
	}
}

func (s *SeedsRun) Run() error {
	err := s.GoalsRepository.DeleteAllGoalCompletions()
	err = s.GoalsRepository.DeleteAllGoals()

	goals := []request.CreateNewGoalRequest{
		{
			Title:                  "Acordar cedo",
			DesiredWeeklyFrequency: 5,
		},
		{
			Title:                  "Me exercitar",
			DesiredWeeklyFrequency: 3,
		},
		{
			Title:                  "Meditar",
			DesiredWeeklyFrequency: 1,
		},
	}

	var createdGoals []db.Goal

	for _, goal := range goals {
		createdGoal, err := s.GoalsRepository.CreateNewGoal(request.CreateNewGoalRequest{
			Title:                  goal.Title,
			DesiredWeeklyFrequency: goal.DesiredWeeklyFrequency,
		})

		if err != nil {
			panic(err)
		}
		
		createdGoals = append(createdGoals, createdGoal)
	}

	err = s.GoalsRepository.CreateNewGoalCompletion(createdGoals[0].ID, time.Now())
	err = s.GoalsRepository.CreateNewGoalCompletion(createdGoals[1].ID, time.Now().AddDate(0, 0, -2))

	fmt.Println("Seeded database!")

	return err
}
