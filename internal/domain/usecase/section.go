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

type SectionUseCase struct {
	repoSection repository.IRepositorySection
	file        storage.IGCImageRepo
}

func NewSectionUseCase(repoSection repository.IRepositorySection, file storage.IGCImageRepo) SectionUseCase {
	return SectionUseCase{
		repoSection,
		file,
	}
}

func (c *SectionUseCase) InsertSection(ctx context.Context, section *entity.Section) (interface{}, int) {
	pathname, err := c.file.SetFile(ctx, section.Cover, "section/section-%s.png")
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	section.Cover = pathname
	newSection, err := c.repoSection.InsertSection(*section)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al insertar la seccion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "seccion insertada exitosamente", newSection), http.StatusOK
}

func (c *SectionUseCase) FindSection(filter map[string]interface{}) (interface{}, int) {
	sections, err := c.repoSection.FindSection(filter)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", sections), http.StatusOK
}

func (c *SectionUseCase) FindSectionOne(idSection int64) (interface{}, int) {
	section, err := c.repoSection.FindSectionOne(idSection)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", section), http.StatusOK
}

func (c *SectionUseCase) DeleteSection(ctx context.Context, idSection int64) (interface{}, int) {
	section, err := c.repoSection.FindSectionOne(idSection)
	if err != nil || section == nil || section.ID == 0 {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	objectName := utils.ExtractObjectName(section.Cover)
	err = c.file.DeleteFile(ctx, "section/%s", objectName)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	err = c.repoSection.DeleteSection(idSection)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar la seccion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "seccion eliminada exitosamente", nil), http.StatusOK
}
