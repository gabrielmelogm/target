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
	GetPendingGoals(firstDayOfWeek time.Time, lastDayOfWeek time.Time) ([]db.GetPendingGoalsRow, error)
	GetGoalCompletionsCountById(goalId string, firstDayOfWeek time.Time, lastDayOfWeek time.Time) (db.GetGoalCompletionsCountByIdRow, error)
	CreateNewGoal(createNewGoalDto request.CreateNewGoalRequest) (db.Goal, error)
	CreateNewGoalCompletion(goalId string) error
	DeleteAllGoals() error
	DeleteAllGoalCompletions() error
}

type GetPendingGoalsResponse struct {
	Id                     string    `json:"id"`
	Title                  string    `json:"title"`
	DesiredWeeklyFrequency int       `json:"desired_weekly_frequency"`
	CreatedAt              time.Time `json:"created_at"`
	CompletionCount        int       `json:"completion_count"`
}

func (g *GoalsRepository) GetPendingGoals(firstDayOfWeek time.Time, lastDayOfWeek time.Time) ([]GetPendingGoalsResponse, error) {
	ctx := context.Background()

	goals, err := g.Queries.GetPendingGoals(ctx, db.GetPendingGoalsParams{
		CreatedAt:   firstDayOfWeek,
		CreatedAt_2: lastDayOfWeek,
	})

	var formatedGoals []GetPendingGoalsResponse

	for _, X := range goals {
		formatedGoals = append(formatedGoals, GetPendingGoalsResponse{
			Id:                     X.ID,
			Title:                  X.Title,
			DesiredWeeklyFrequency: int(X.DesiredWeeklyFrequency),
			CreatedAt:              X.CreatedAt,
			CompletionCount:        int(X.CompletionCount.Int64),
		})
	}

	return formatedGoals, err
}

func (g *GoalsRepository) GetGoalCompletionsCountById(goalId string, firstDayOfWeek time.Time, lastDayOfWeek time.Time) (db.GetGoalCompletionsCountByIdRow, error) {
	ctx := context.Background()

	completionsGoal, err := g.Queries.GetGoalCompletionsCountById(ctx, db.GetGoalCompletionsCountByIdParams{
		CreatedAt:   firstDayOfWeek,
		CreatedAt_2: lastDayOfWeek,
		ID:          goalId,
	})

	return completionsGoal, err
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
