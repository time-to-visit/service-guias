package entry

import (
	"net/http"
	"service-user/internal/domain/entity"
	"service-user/internal/domain/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type activitiesEntry struct {
	activitiesCaseuse usecase.ActivitiesUseCase
}

func NewActivitiesEntry(activitiesCaseuse usecase.ActivitiesUseCase) *activitiesEntry {
	return &activitiesEntry{
		activitiesCaseuse,
	}
}

func (c *activitiesEntry) RegisterActivities(context echo.Context) error {
	activities := context.Get("activities").(*entity.Activities)
	response, status := c.activitiesCaseuse.RegisterActivities(context.Request().Context(), activities)
	return context.JSON(status, response)
}

func (c *activitiesEntry) FindActivities(context echo.Context) error {
	id := context.Param("ID")
	idCategory, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := c.activitiesCaseuse.FindActivities(int64(idCategory))
	return context.JSON(status, response)
}

func (c *activitiesEntry) FindActivitiesOne(context echo.Context) error {
	id := context.Param("ID")
	idCategory, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := c.activitiesCaseuse.FindActivitiesOne(int64(idCategory))
	return context.JSON(status, response)

}

func (c *activitiesEntry) DeleteActivities(context echo.Context) error {
	id := context.Param("ID")
	idActivities, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := c.activitiesCaseuse.DeleteActivities(context.Request().Context(), int64(idActivities))
	return context.JSON(status, response)
}
