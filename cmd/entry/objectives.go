package entry

import (
	"net/http"
	"service-user/internal/domain/entity"
	"service-user/internal/domain/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type objectiveEntry struct {
	objectiveCaseuse usecase.ObjectivesUseCase
}

func NewOjectiveEntry(objectiveCaseuse usecase.ObjectivesUseCase) *objectiveEntry {
	return &objectiveEntry{
		objectiveCaseuse,
	}
}

func (c *objectiveEntry) RegisterObjective(context echo.Context) error {
	objectives := context.Get("objectives").(*entity.Objectives)
	response, status := c.objectiveCaseuse.RegisterObjective(objectives)
	return context.JSON(status, response)
}

func (c *objectiveEntry) DeleteObjective(context echo.Context) error {
	id := context.Param("ID")
	idObjetives, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := c.objectiveCaseuse.DeleteObjective(int64(idObjetives))
	return context.JSON(status, response)
}

func (c *objectiveEntry) FindObjectiveByLevels(context echo.Context) error {
	id := context.Param("ID")
	idLevels, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := c.objectiveCaseuse.FindObjectiveByLevels(int64(idLevels))
	return context.JSON(status, response)
}

func (c *objectiveEntry) FindObjectiveOne(context echo.Context) error {
	id := context.Param("ID")
	idObjetives, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := c.objectiveCaseuse.FindObjectiveOne(int64(idObjetives))
	return context.JSON(status, response)
}
