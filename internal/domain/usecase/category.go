package usecase

import (
	"context"
	"net/http"
	"service-user/internal/domain/entity"
	objectValues "service-user/internal/domain/object_values"
	"service-user/internal/domain/repository"
	"service-user/internal/infra/storage"
	"service-user/internal/utils"
)

type CategoriesUseCase struct {
	repoCategories repository.IRepositoryCategories
	file           storage.IGCImageRepo
}

func NewCateogriesUseCase(repoCategories repository.IRepositoryCategories, file storage.IGCImageRepo) CategoriesUseCase {
	return CategoriesUseCase{
		repoCategories,
		file,
	}
}

func (c *CategoriesUseCase) RegisterCategory(ctx context.Context, category *entity.Categories) (interface{}, int) {
	pathname, err := c.file.SetFile(ctx, category.Cover, "category-guides/cat-%s.png")
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	category.Cover = pathname
	newCategories, err := c.repoCategories.InsertCategory(*category)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al insertar la categoria", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "categoria insertada exitosamente", newCategories), http.StatusOK
}

func (c *CategoriesUseCase) FindCategory(idSections int64) (interface{}, int) {
	categories, err := c.repoCategories.FindCategory(idSections)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", categories), http.StatusOK
}

func (c *CategoriesUseCase) FindCategoryOne(idCategory int64) (interface{}, int) {
	category, err := c.repoCategories.FindCategoryOne(idCategory)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", category), http.StatusOK
}

func (c *CategoriesUseCase) DeleteCategory(ctx context.Context, idCategory int64) (interface{}, int) {
	category, err := c.repoCategories.FindCategoryOne(idCategory)
	if err != nil || category == nil || category.ID == 0 {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "no se encontro la categoria", nil), http.StatusBadRequest
	}
	objectName := utils.ExtractObjectName(category.Cover)
	err = c.file.DeleteFile(ctx, "category-guides/%s", objectName)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar la guias", nil), http.StatusBadRequest
	}
	err = c.repoCategories.DeleteCategory(idCategory)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar la categoria", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "categoria eliminada exitosamente", nil), http.StatusOK
}
