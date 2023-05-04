package entry

import (
	"net/http"
	"service-user/internal/domain/entity"
	"service-user/internal/domain/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type sectionEntry struct {
	sectionCaseuse usecase.SectionUseCase
}

func NewSectionEntry(sectionCaseuse usecase.SectionUseCase) *sectionEntry {
	return &sectionEntry{
		sectionCaseuse,
	}
}

func (c *sectionEntry) InsertSection(context echo.Context) error {
	section := context.Get("section").(*entity.Section)
	response, status := c.sectionCaseuse.InsertSection(context.Request().Context(), section)
	return context.JSON(status, response)
}

func (c *sectionEntry) FindSection(context echo.Context) error {
	filter := make(map[string]interface{})
	response, status := c.sectionCaseuse.FindSection(filter)
	return context.JSON(status, response)
}

func (c *sectionEntry) FindSectionOne(context echo.Context) error {
	id := context.Param("ID")
	idSection, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := c.sectionCaseuse.FindSectionOne(int64(idSection))
	return context.JSON(status, response)
}

func (c *sectionEntry) DeleteSection(context echo.Context) error {
	id := context.Param("ID")
	idSection, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := c.sectionCaseuse.DeleteSection(context.Request().Context(), int64(idSection))
	return context.JSON(status, response)
}
