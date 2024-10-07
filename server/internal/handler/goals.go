package handler

import (
	"fmt"
	"net/http"
	"target/internal/request"
	"target/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type GoalsHandler struct {
	GoalsService service.GoalsService
}

func NewGoalHandler(goalsService service.GoalsService) *GoalsHandler {
	return &GoalsHandler{
		GoalsService: goalsService,
	}
}

func (g *GoalsHandler) CreateNewGoal(c echo.Context) error {
	var data request.CreateNewGoalRequest

	err := c.Bind(&data)

	validate := validator.New()

	err = validate.Struct(&data)

	if err := validate.Struct(data); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("Field '%s' failed validation with tag '%s'", err.Field(), err.Tag())
		}
		return c.JSON(http.StatusBadRequest, echo.Map{"errors": errors})
	}

	createdGoal, err := g.GoalsService.CreateNewGoal(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, createdGoal)
}
