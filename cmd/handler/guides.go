package handler

import (
	"service-user/cmd/entry"
	"service-user/internal/domain/usecase"
	"service-user/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerGuide(e *echo.Echo, guideUseCase usecase.GuidesUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	guideEntry := entry.NewGuideEntry(guideUseCase)
	e.DELETE("/guides/guides/:ID", guideEntry.DeleteGuides, auth)
	e.POST("/guides/guides", guideEntry.RegisterGuides, auth, validator.ValidateGuides)
	e.GET("/guides/guides", guideEntry.FindGuides, auth)
	e.GET("/guides/guides/:ID", guideEntry.FindGuidesOne, auth)
	return e
}
