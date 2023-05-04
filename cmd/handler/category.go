package handler

import (
	"service-user/cmd/entry"
	"service-user/internal/domain/usecase"
	"service-user/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerCategories(e *echo.Echo, categoriesUseCase usecase.CategoriesUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	categoriesEntry := entry.NewCategoryEntry(categoriesUseCase)
	e.POST("/categories", categoriesEntry.RegisterCategory, auth, validator.ValidateCategories)
	e.DELETE("/categories/:ID", categoriesEntry.DeleteCategory, auth)
	e.GET("/categories", categoriesEntry.FindCategory, auth)
	e.GET("/categories/:ID", categoriesEntry.FindCategoryOne, auth)
	return e
}
