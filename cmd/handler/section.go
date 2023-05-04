package handler

import (
	"service-user/cmd/entry"
	"service-user/internal/domain/usecase"
	"service-user/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerSection(e *echo.Echo, sectionUseCase usecase.SectionUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	sectionEntry := entry.NewSectionEntry(sectionUseCase)
	e.POST("/section", sectionEntry.InsertSection, auth, validator.ValidateSections)
	e.DELETE("/section/:ID", sectionEntry.DeleteSection, auth)
	e.GET("/section", sectionEntry.FindSection, auth)
	e.GET("/section/:ID", sectionEntry.FindSectionOne, auth)
	return e
}
