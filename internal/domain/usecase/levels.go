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

type LevelsUseCase struct {
	repoLevels repository.IRepositoryLevels
	file       storage.IGCImageRepo
}

func NewLevelsUseCase(repoLevels repository.IRepositoryLevels, file storage.IGCImageRepo) LevelsUseCase {
	return LevelsUseCase{
		repoLevels,
		file,
	}
}

func (c *LevelsUseCase) InsertLevels(ctx context.Context, levels entity.Levels) (interface{}, int) {
	pathname, err := c.file.SetFile(ctx, levels.Cover, "levels-guides/levels-%s.png")
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	levels.Cover = pathname
	newLevel, err := c.repoLevels.InsertLevels(levels)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al insertar new level", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "level insertado exitosamente", newLevel), http.StatusOK

}

func (c *LevelsUseCase) FindLevelsByCategory(idCategory int64) (interface{}, int) {
	levels, err := c.repoLevels.FindLevelsByCategory(idCategory)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", levels), http.StatusOK

}

func (c *LevelsUseCase) FindLevelsOne(idLevels int64) (interface{}, int) {
	level, err := c.repoLevels.FindLevelsOne(idLevels)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", level), http.StatusOK

}

func (c *LevelsUseCase) DeleteLevels(ctx context.Context, idLevels int64) (interface{}, int) {
	level, err := c.repoLevels.FindLevelsOne(idLevels)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	objectName := utils.ExtractObjectName(level.Cover)
	err = c.file.DeleteFile(ctx, "levels-guides/%s", objectName)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar la guias", nil), http.StatusBadRequest
	}
	err = c.repoLevels.DeleteLevels(idLevels)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar el contenido del nivel", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "error", "nivel eliminado sastifactoriamente", nil), http.StatusOK
}
