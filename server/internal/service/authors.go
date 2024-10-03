package service

import (
	"database/sql"
	"errors"
	"target/internal/entity"
	"target/internal/repository"
	"target/internal/request"
)

type AuthorsServiceInterface interface {
	CreateNewAuthor(arg request.CreateNewAuthorRequest)
}

type AuthorsService struct {
	AuthorsRepository repository.AuthorsRepository
}

func NewAuthorsService(AuthorsRepository repository.AuthorsRepository) *AuthorsService {
	return &AuthorsService{
		AuthorsRepository: AuthorsRepository,
	}
}

func (a *AuthorsService) GetAuthorByName(name string) (entity.Author, error) {
	findedAuthor, err := a.AuthorsRepository.GetAuthorByName(name)

	if findedAuthor.Name != "" {
		return entity.Author{}, errors.New("User with this name already exists")	
	}

	var author entity.Author = entity.Author{
		ID:   findedAuthor.ID,
		Name: findedAuthor.Name,
		Bio:  findedAuthor.Bio.String,
	}

	return author, err
}

func (a *AuthorsService) CreateNewAuthor(arg request.CreateNewAuthorRequest) (entity.Author, error) {
	findedAuthor, err := a.GetAuthorByName(arg.Name)

	if !errors.Is(err, sql.ErrNoRows) {
		return findedAuthor, err
	}

	createdAuthor, err := a.AuthorsRepository.CreateNewAuthor(request.CreateNewAuthorRequest{
		Name: arg.Name,
		Bio:  arg.Bio,
	})

	var author entity.Author = entity.Author{
		ID:   createdAuthor.ID,
		Name: createdAuthor.Name,
		Bio:  createdAuthor.Bio.String,
	}

	return author, err
}
