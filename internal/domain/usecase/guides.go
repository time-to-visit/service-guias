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

type GuidesUseCase struct {
	repoGuides repository.IRepositoryGuides
	file       storage.IGCImageRepo
}

func NewGuidesUseCase(repoGuides repository.IRepositoryGuides, file storage.IGCImageRepo) GuidesUseCase {
	return GuidesUseCase{
		repoGuides,
		file,
	}
}

func (c *GuidesUseCase) RegisterGuides(ctx context.Context, guide *entity.Guides) (interface{}, int) {
	pathname, err := c.file.SetFile(ctx, guide.Cover, "guides/guides-%s.png")
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	guide.Cover = pathname
	newGuide, err := c.repoGuides.InsertGuides(*guide)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al insertar la seccion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "guia insertada exitosamente", newGuide), http.StatusOK
}

func (c *GuidesUseCase) FindGuides(filter map[string]interface{}) (interface{}, int) {
	guides, err := c.repoGuides.FindGuides(filter)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", guides), http.StatusOK
}

func (c *GuidesUseCase) FindGuidesOne(idGuides int64) (interface{}, int) {
	guide, err := c.repoGuides.FindGuidesOne(idGuides)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", guide), http.StatusOK
}

func (c *GuidesUseCase) DeleteGuides(ctx context.Context, idGuides int64) (interface{}, int) {
	guide, err := c.repoGuides.FindGuidesOne(idGuides)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	objectName := utils.ExtractObjectName(guide.Cover)
	err = c.file.DeleteFile(ctx, "guides/%s", objectName)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar la guias", nil), http.StatusBadRequest
	}
	err = c.repoGuides.DeleteGuides(idGuides)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar la guia", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "guia eliminada exitosamente", nil), http.StatusOK
}
