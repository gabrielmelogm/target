package seeds

import "target/internal/repository"

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

	return err
}
