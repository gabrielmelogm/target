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

	api := e.Group("/api")

	api.GET("/healthcheck", goal.GetHealthCheck)
	api.GET("/summary", goal.GetWeekSummary)
	api.GET("/pending-goals", goal.GetPendingGoals)
	api.POST("/completions", goal.CreateNewGoalCompletionById)
	api.POST("/goal", goal.CreateNewGoal)

	e.Logger.Fatal(e.Start(":3333"))
}
