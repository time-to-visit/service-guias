package entry

import (
	"net/http"
	"service-user/internal/domain/entity"
	"service-user/internal/domain/usecase"
	"service-user/internal/domain/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type guideEntry struct {
	guideCaseuse usecase.GuidesUseCase
}

func NewGuideEntry(guideCaseuse usecase.GuidesUseCase) *guideEntry {
	return &guideEntry{
		guideCaseuse,
	}
}

func (c *guideEntry) RegisterGuides(context echo.Context) error {
	guides := context.Get("guides").(*entity.Guides)
	response, status := c.guideCaseuse.RegisterGuides(context.Request().Context(), guides)
	return context.JSON(status, response)
}

func (c *guideEntry) FindGuides(context echo.Context) error {
	filter := utils.UrlValuesToMap(context.QueryParams())
	response, status := c.guideCaseuse.FindGuides(filter)
	return context.JSON(status, response)
}

func (c *guideEntry) FindGuidesOne(context echo.Context) error {
	id := context.Param("ID")
	idGuides, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := c.guideCaseuse.FindGuidesOne(int64(idGuides))
	return context.JSON(status, response)
}

func (c *guideEntry) DeleteGuides(context echo.Context) error {
	id := context.Param("ID")
	idGuides, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := c.guideCaseuse.DeleteGuides(context.Request().Context(), int64(idGuides))
	return context.JSON(status, response)
}
