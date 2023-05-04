package usecase

import (
	"net/http"
	"service-user/internal/domain/entity"
	objectValues "service-user/internal/domain/object_values"
	"service-user/internal/domain/repository"
)

type ObjectivesUseCase struct {
	repoObjectives repository.IRepositoryObjectives
}

func NewObjectivesUseCase(repoObjectives repository.IRepositoryObjectives) ObjectivesUseCase {
	return ObjectivesUseCase{
		repoObjectives,
	}
}

func (c *ObjectivesUseCase) RegisterObjective(objetives *entity.Objectives) (interface{}, int) {
	newObjective, err := c.repoObjectives.InsertObjective(*objetives)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al insertar new objectives", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "objectivo insertado exitosamente", newObjective), http.StatusOK

}

func (c *ObjectivesUseCase) DeleteObjective(idObjective int64) (interface{}, int) {
	err := c.repoObjectives.DeleteObjective(idObjective)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar el objectivo", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "objectivo eliminado exitosamente", nil), http.StatusOK

}

func (c *ObjectivesUseCase) FindObjectiveByLevels(idLevel int64) (interface{}, int) {
	objectives, err := c.repoObjectives.FindObjectiveByLevel(idLevel)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", objectives), http.StatusOK
}

func (c *ObjectivesUseCase) FindObjectiveOne(idObective int64) (interface{}, int) {
	objective, err := c.repoObjectives.FindObjectiveOne(idObective)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", objective), http.StatusOK
}
