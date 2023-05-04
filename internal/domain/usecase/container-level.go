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

type ContainerLevelsUseCase struct {
	repoContainerLevels repository.IRepositoryContainerLevels
	file                storage.IGCImageRepo
}

func NewContainerLevelsUseCase(repoContainerLevels repository.IRepositoryContainerLevels, file storage.IGCImageRepo) ContainerLevelsUseCase {
	return ContainerLevelsUseCase{
		repoContainerLevels,
		file,
	}
}

func (c *ContainerLevelsUseCase) InsertContainerLevels(ctx context.Context, container *entity.ContainerLevels) (interface{}, int) {
	if container.Type == "IMAGE" {
		pathname, err := c.file.SetFile(ctx, container.Container, "container-level/container-%s.png")
		if err != nil {
			return objectValues.NewResponseWithData(http.StatusBadRequest,
				"error",
				"hubo un problema con storage",
				nil,
			), http.StatusBadRequest
		}
		container.Container = pathname
	}

	container, err := c.repoContainerLevels.InsertContainerLevels(*container)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al insertar la seccion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "guia insertada exitosamente", container), http.StatusOK
}

func (c *ContainerLevelsUseCase) FindContainerLevels(idLevels int64) (interface{}, int) {
	container, err := c.repoContainerLevels.FindContainerLevels(idLevels)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", container), http.StatusOK
}

func (c *ContainerLevelsUseCase) FindContainerLevelsOne(idContainerLevels int64) (interface{}, int) {
	container, err := c.repoContainerLevels.FindContainerLevelsOne(idContainerLevels)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", container), http.StatusOK

}

func (c *ContainerLevelsUseCase) DeleteContainerLevels(ctx context.Context, idContainerLevels int64) (interface{}, int) {
	container, err := c.repoContainerLevels.FindContainerLevelsOne(idContainerLevels)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	if container.Type == "IMAGE" {
		objectName := utils.ExtractObjectName(container.Container)
		err = c.file.DeleteFile(ctx, "container-level/%s", objectName)
		if err != nil {
			return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar el contenido del nivel", nil), http.StatusBadRequest
		}
	}
	err = c.repoContainerLevels.DeleteContainerLevels(idContainerLevels)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar el contenido del nivel", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "contenido del nivel eliminado sastifactoriamente", nil), http.StatusBadRequest
}
