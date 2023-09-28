package handler

import (
	"service-user/cmd/entry"
	"service-user/internal/domain/usecase"
	"service-user/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerCategories(e *echo.Echo, categoriesUseCase usecase.CategoriesUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	categoriesEntry := entry.NewCategoryEntry(categoriesUseCase)
	e.POST("/guides/categories", categoriesEntry.RegisterCategory, auth, validator.ValidateCategories)
	e.DELETE("/guides/categories/:ID", categoriesEntry.DeleteCategory, auth)
	e.GET("/guides/categories/section/:ID", categoriesEntry.FindCategory, auth)
	e.GET("/guides/categories/:ID", categoriesEntry.FindCategoryOne, auth)
	return e
}
