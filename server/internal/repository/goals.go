package repository

import (
	"context"
	"database/sql"
	db "target/db/sqlc"
	"target/internal/request"
	"time"

	"github.com/google/uuid"
)

type GoalsRepository struct {
	Queries *db.Queries
	Conn    *sql.DB
}

func NewGoalsRepository(conn *sql.DB) *GoalsRepository {
	queries := db.New(conn)
	return &GoalsRepository{
		Queries: queries,
		Conn:    conn,
	}
}

type GoalsRepositoryInterface interface {
	CreateNewGoal(createNewGoalDto request.CreateNewGoalRequest) (db.Goal, error)
	CreateNewGoalCompletion(goalId string) error
	DeleteAllGoals() error
	DeleteAllGoalCompletions() error
}

func (g *GoalsRepository) CreateNewGoal(createNewGoalDto request.CreateNewGoalRequest) (db.Goal, error) {
	ctx := context.Background()

	createdGoal, err := g.Queries.CreateNewGoal(ctx, db.CreateNewGoalParams{
		ID:                     uuid.New().String(),
		Title:                  createNewGoalDto.Title,
		DesiredWeeklyFrequency: int32(createNewGoalDto.DesiredWeeklyFrequency),
	})

	return createdGoal, err
}

func (g *GoalsRepository) CreateNewGoalCompletion(goalId string, createdAt time.Time) error {
	ctx := context.Background()

	completeData := createdAt

	if createdAt.IsZero() {
		completeData = time.Now()
	}

	err := g.Queries.CreateNewGoalCompletion(ctx, db.CreateNewGoalCompletionParams{
		ID:        uuid.New().String(),
		GoalID:    goalId,
		CreatedAt: completeData,
	})
	return err
}

func (g *GoalsRepository) DeleteAllGoals() error {
	ctx := context.Background()

	err := g.Queries.DeleteAllGoals(ctx)

	return err
}

func (g *GoalsRepository) DeleteAllGoalCompletions() error {
	ctx := context.Background()

	err := g.Queries.DeleteAllGoalCompletions(ctx)

	return err
}
