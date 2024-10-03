package cmd

import (
	"github.com/labstack/echo/v4"
)

func Server() {
	e := echo.New()

	e.Logger.Fatal(e.Start(":3333"))
}
