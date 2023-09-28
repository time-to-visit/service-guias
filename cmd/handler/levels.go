package handler

import (
	"service-user/cmd/entry"
	"service-user/internal/domain/usecase"
	"service-user/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerLevel(e *echo.Echo, levelUseCase usecase.LevelsUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	levelEntry := entry.NewLevelEntry(levelUseCase)
	e.POST("/guides/level", levelEntry.RegisterLevels, auth, validator.ValidateLevels)
	e.DELETE("/guides/level/:ID", levelEntry.DeleteLevels, auth)
	e.GET("/guides/level/category/:ID", levelEntry.FindLevelsByCategory, auth)
	e.GET("/guides/level/:ID", levelEntry.FindLevelsOne, auth)
	return e
}
