package handler

import (
	"service-user/cmd/entry"
	"service-user/internal/domain/usecase"
	"service-user/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerActivities(e *echo.Echo, activitiesUseCase usecase.ActivitiesUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	activitiesEntry := entry.NewActivitiesEntry(activitiesUseCase)
	e.DELETE("/activities/:ID", activitiesEntry.DeleteActivities, auth)
	e.POST("/activities", activitiesEntry.RegisterActivities, auth, validator.ValidateActivities)
	e.GET("/activities/:ID", activitiesEntry.FindActivities, auth)
	e.GET("/activities/one/:ID", activitiesEntry.FindActivitiesOne, auth)
	return e
}
