package seeds

import (
	"target/internal/repository"
	"target/internal/request"
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

	createdGoal, err := s.GoalsRepository.CreateNewGoal(request.CreateNewGoalRequest{
		Title:                  "John Doe",
		DesiredWeeklyFrequency: 4,
	})

	err = s.GoalsRepository.CreateNewGoalCompletion(createdGoal.ID)

	if err != nil {
		panic(err)
	}

	return err
}
