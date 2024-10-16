package cmd

import (
	"target/internal/handler"

	"github.com/labstack/echo/v4"
)

func Server(goal *handler.GoalsHandler) {
	e := echo.New()

	e.GET("/summary", goal.GetWeekSummary)
	e.GET("/pending-goals", goal.GetPendingGoals)
	e.POST("/completions", goal.CreateNewGoalCompletionById)
	e.POST("/goal", goal.CreateNewGoal)

	e.Logger.Fatal(e.Start(":3333"))
}
