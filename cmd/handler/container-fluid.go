package handler

import (
	"service-user/cmd/entry"
	"service-user/internal/domain/usecase"
	"service-user/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerContainerLevel(e *echo.Echo, containerLevelUseCase usecase.ContainerLevelsUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	containerLevelEntry := entry.NewContainerLevelEntry(containerLevelUseCase)
	e.POST("/guides/container-level", containerLevelEntry.InsertContainerLevels, auth, validator.ValidateContainerLevels)
	e.DELETE("/guides/container-level/:ID", containerLevelEntry.DeleteContainerLevels, auth)
	e.GET("/guides/container-level/:ID", containerLevelEntry.FindContainerLevelsOne, auth)
	e.GET("/guides/container-level/level/:ID", containerLevelEntry.FindContainerLevels, auth)

	return e
}
