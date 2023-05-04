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

type ActivitiesUseCase struct {
	repoActivities repository.IRepositoryActivities
	file           storage.IGCImageRepo
}

func NewActivitiesUseCase(repoActivities repository.IRepositoryActivities, file storage.IGCImageRepo) ActivitiesUseCase {
	return ActivitiesUseCase{
		repoActivities,
		file,
	}
}

func (r *ActivitiesUseCase) RegisterActivities(ctx context.Context, activities *entity.Activities) (interface{}, int) {
	pathname, err := r.file.SetFile(ctx, activities.Image, "activities/act-%s.png")
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest,
			"error",
			"hubo un problema con storage",
			nil,
		), http.StatusBadRequest
	}
	activities.Image = pathname
	newActivities, err := r.repoActivities.RegisterActivities(*activities)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al insertar la actividad", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "actividad insertada exitosamente", newActivities), http.StatusOK
}

func (r *ActivitiesUseCase) FindActivities(idSites int64) (interface{}, int) {
	activties, err := r.repoActivities.FindActivities(idSites)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", activties), http.StatusOK
}

func (r *ActivitiesUseCase) FindActivitiesOne(idActivities int64) (interface{}, int) {
	activity, err := r.repoActivities.FindActivitiesOne(idActivities)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", activity), http.StatusOK
}

func (r *ActivitiesUseCase) DeleteActivities(ctx context.Context, idActivities int64) (interface{}, int) {
	activity, err := r.repoActivities.FindActivitiesOne(idActivities)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	objectName := utils.ExtractObjectName(activity.Image)
	err = r.file.DeleteFile(ctx, "container-level/%s", objectName)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar el contenido del nivel", nil), http.StatusBadRequest
	}
	err = r.repoActivities.DeleteActivities(idActivities)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar la actividad", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "actividad eliminado sastifactoriamente", nil), http.StatusBadRequest

}
