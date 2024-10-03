package cmd

import (
	"target/internal/handler"

	"github.com/labstack/echo/v4"
)

func Server(author *handler.AuthorHandler) {
	e := echo.New()

	e.POST("/author", author.CreateAuthor)

	e.Logger.Fatal(e.Start(":3333"))
}
