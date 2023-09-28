package handler

import (
	"service-user/cmd/entry"
	"service-user/internal/domain/usecase"
	"service-user/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerObjetive(e *echo.Echo, objectiveUseCase usecase.ObjectivesUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	objectiveEntry := entry.NewOjectiveEntry(objectiveUseCase)
	e.POST("/guides/objectives", objectiveEntry.RegisterObjective, auth, validator.ValidateObjectives)
	e.DELETE("/guides/objectives/:ID", objectiveEntry.DeleteObjective, auth)
	e.GET("/guides/objectives/level/:ID", objectiveEntry.FindObjectiveByLevels, auth)
	e.GET("/guides/objectives/:ID", objectiveEntry.FindObjectiveOne, auth)
	return e
}
