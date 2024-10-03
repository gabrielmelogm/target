package handler

import (
	"net/http"
	"target/internal/request"
	"target/internal/service"

	"github.com/labstack/echo/v4"
)

type AuthorHandler struct {
	AuthorsService service.AuthorsService
}

func NewAuthorHandler(
	AuthorsService service.AuthorsService,
) *AuthorHandler {
	return &AuthorHandler{
		AuthorsService: AuthorsService,
	}
}

func (a *AuthorHandler) CreateAuthor(c echo.Context) error {
	var data request.CreateNewAuthorRequest

	err := c.Bind(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	author, err := a.AuthorsService.CreateNewAuthor(request.CreateNewAuthorRequest{
		Name: data.Name,
		Bio:  data.Bio,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, author)
}
