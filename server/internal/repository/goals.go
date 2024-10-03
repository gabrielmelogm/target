package repository

import (
	"context"
	"database/sql"
	db "target/db/sqlc"
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
	DeleteAllGoals() error
	DeleteAllGoalCompletions() error
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
