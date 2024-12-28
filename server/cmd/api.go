package cmd

import (
	"target/internal/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Server(goal *handler.GoalsHandler) {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           86400,
		AllowCredentials: true,
	}))

	e.Group("/api")

	e.GET("/summary", goal.GetWeekSummary)
	e.GET("/pending-goals", goal.GetPendingGoals)
	e.POST("/completions", goal.CreateNewGoalCompletionById)
	e.POST("/goal", goal.CreateNewGoal)

	e.Logger.Fatal(e.Start(":3333"))
}
