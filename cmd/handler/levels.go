package handler

import (
	"service-user/cmd/entry"
	"service-user/internal/domain/usecase"
	"service-user/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerLevel(e *echo.Echo, levelUseCase usecase.LevelsUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	levelEntry := entry.NewLevelEntry(levelUseCase)
	e.POST("/level", levelEntry.RegisterLevels, auth, validator.ValidateLevels)
	e.DELETE("/level/:ID", levelEntry.DeleteLevels, auth)
	e.GET("/level/category/:ID", levelEntry.FindLevelsByCategory, auth)
	e.GET("/level/:ID", levelEntry.FindLevelsOne, auth)
	return e
}
