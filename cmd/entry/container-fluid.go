package entry

import (
	"net/http"
	"service-user/internal/domain/entity"
	"service-user/internal/domain/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type containerLevelEntry struct {
	containerLevelCaseuse usecase.ContainerLevelsUseCase
}

func NewContainerLevelEntry(containerLevelCaseuse usecase.ContainerLevelsUseCase) *containerLevelEntry {
	return &containerLevelEntry{
		containerLevelCaseuse,
	}
}

func (c *containerLevelEntry) InsertContainerLevels(context echo.Context) error {
	containerLevels := context.Get("containerLevels").(*entity.ContainerLevels)
	response, status := c.containerLevelCaseuse.InsertContainerLevels(context.Request().Context(), containerLevels)
	return context.JSON(status, response)
}
func (c *containerLevelEntry) FindContainerLevels(context echo.Context) error {
	id := context.Param("ID")
	idLevel, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := c.containerLevelCaseuse.FindContainerLevels(int64(idLevel))
	return context.JSON(status, response)
}
func (c *containerLevelEntry) FindContainerLevelsOne(context echo.Context) error {
	id := context.Param("ID")
	idContainerLevel, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := c.containerLevelCaseuse.FindContainerLevelsOne(int64(idContainerLevel))
	return context.JSON(status, response)
}
func (c *containerLevelEntry) DeleteContainerLevels(context echo.Context) error {
	id := context.Param("ID")
	idContainerLevel, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := c.containerLevelCaseuse.DeleteContainerLevels(context.Request().Context(), int64(idContainerLevel))
	return context.JSON(status, response)
}
