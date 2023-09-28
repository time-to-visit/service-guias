package handler

import (
	"service-user/cmd/entry"
	"service-user/internal/domain/usecase"
	"service-user/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerSection(e *echo.Echo, sectionUseCase usecase.SectionUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	sectionEntry := entry.NewSectionEntry(sectionUseCase)
	e.POST("/guides/section", sectionEntry.InsertSection, auth, validator.ValidateSections)
	e.DELETE("/guides/section/:ID", sectionEntry.DeleteSection, auth)
	e.GET("/guides/section", sectionEntry.FindSection, auth)
	e.GET("/guides/section/:ID", sectionEntry.FindSectionOne, auth)
	return e
}
