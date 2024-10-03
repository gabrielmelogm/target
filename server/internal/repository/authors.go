package repository

import (
	"context"
	"database/sql"
	db "target/db/sqlc"
	"target/internal/request"
)

type AuthorInterface interface {
	CreateNewAuthor(ctx context.Context, arg db.CreateAuthorParams) (db.Author, error)
}

type AuthorsRepository struct {
	Queries *db.Queries
	Conn    *sql.DB
}

func NewAuthorsRepository(conn *sql.DB) *AuthorsRepository {
	queries := db.New(conn)
	return &AuthorsRepository{
		Queries: queries,
		Conn:    conn,
	}
}

func (a *AuthorsRepository) CreateNewAuthor(arg request.CreateNewAuthorRequest) (db.Author, error) {
	ctx := context.Background()
	createdAuthor, err := a.Queries.CreateAuthor(ctx, db.CreateAuthorParams{
		Name: arg.Name,
		Bio: sql.NullString{
			String: arg.Bio,
			Valid:  true,
		},
	})

	return createdAuthor, err
}
