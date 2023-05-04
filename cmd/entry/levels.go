package entry

import (
	"net/http"
	"service-user/internal/domain/entity"
	"service-user/internal/domain/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type levelEntry struct {
	levelCaseuse usecase.LevelsUseCase
}

func NewLevelEntry(levelCaseuse usecase.LevelsUseCase) *levelEntry {
	return &levelEntry{
		levelCaseuse,
	}
}

func (c *levelEntry) RegisterLevels(context echo.Context) error {
	levels := context.Get("levels").(*entity.Levels)
	response, status := c.levelCaseuse.InsertLevels(context.Request().Context(), *levels)
	return context.JSON(status, response)
}

func (c *levelEntry) FindLevelsByCategory(context echo.Context) error {
	id := context.Param("ID")
	idCategory, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := c.levelCaseuse.FindLevelsByCategory(int64(idCategory))
	return context.JSON(status, response)
}

func (c *levelEntry) FindLevelsOne(context echo.Context) error {
	id := context.Param("ID")
	idLevel, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := c.levelCaseuse.FindLevelsOne(int64(idLevel))
	return context.JSON(status, response)
}

func (c *levelEntry) DeleteLevels(context echo.Context) error {
	id := context.Param("ID")
	idLevel, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := c.levelCaseuse.DeleteLevels(context.Request().Context(), int64(idLevel))
	return context.JSON(status, response)
}
