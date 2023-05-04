package entry

import (
	"net/http"
	"service-user/internal/domain/entity"
	"service-user/internal/domain/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type categoryEntry struct {
	categoryCaseuse usecase.CategoriesUseCase
}

func NewCategoryEntry(categoryCaseuse usecase.CategoriesUseCase) *categoryEntry {
	return &categoryEntry{
		categoryCaseuse,
	}
}

func (r *categoryEntry) RegisterCategory(context echo.Context) error {
	categories := context.Get("categories").(*entity.Categories)
	response, status := r.categoryCaseuse.RegisterCategory(context.Request().Context(), categories)
	return context.JSON(status, response)
}

func (r *categoryEntry) FindCategory(context echo.Context) error {
	id := context.Param("ID")
	idSection, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := r.categoryCaseuse.FindCategory(int64(idSection))
	return context.JSON(status, response)
}

func (r *categoryEntry) FindCategoryOne(context echo.Context) error {
	id := context.Param("ID")
	idCategory, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := r.categoryCaseuse.FindCategoryOne(int64(idCategory))
	return context.JSON(status, response)
}

func (r *categoryEntry) DeleteCategory(context echo.Context) error {
	id := context.Param("ID")
	idCategory, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := r.categoryCaseuse.DeleteCategory(context.Request().Context(), int64(idCategory))
	return context.JSON(status, response)
}
