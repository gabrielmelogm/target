package cmd

import (
	"target/internal/handler"

	"github.com/labstack/echo/v4"
)

func Server(goal *handler.GoalsHandler) {
	e := echo.New()

	e.POST("/goal", goal.CreateNewGoal)

	e.Logger.Fatal(e.Start(":3333"))
}
