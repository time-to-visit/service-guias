package usecase_test

import (
	"context"
	"errors"
	"net/http"
	"service-user/internal/domain/entity"
	"service-user/internal/domain/usecase"
	"service-user/mocks"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func Test_RegisterActivitiesOk(t *testing.T) {
	activites := entity.Activities{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryActivities(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	repoAct.On("RegisterActivities", mock.Anything).Return(&activites, nil)
	usecaseAct := usecase.NewActivitiesUseCase(repoAct, objectFile)
	_, status := usecaseAct.RegisterActivities(context.Background(), &activites)
	assert.Equal(t, status, http.StatusOK)
}

func Test_RegisterActivitiesErrSetFile(t *testing.T) {
	activites := entity.Activities{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryActivities(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New(""))
	usecaseAct := usecase.NewActivitiesUseCase(repoAct, objectFile)
	_, status := usecaseAct.RegisterActivities(context.Background(), &activites)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_RegisterActivitiesErrRegister(t *testing.T) {
	activites := entity.Activities{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryActivities(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	repoAct.On("RegisterActivities", mock.Anything).Return(&activites, errors.New(""))
	usecaseAct := usecase.NewActivitiesUseCase(repoAct, objectFile)
	_, status := usecaseAct.RegisterActivities(context.Background(), &activites)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_FindActivitiesOk(t *testing.T) {
	activites := []entity.Activities{
		{
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	repoAct := mocks.NewIRepositoryActivities(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindActivities", mock.Anything).Return(&activites, nil)
	usecaseAct := usecase.NewActivitiesUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindActivities(1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_FindActivitiesErr(t *testing.T) {
	activites := []entity.Activities{
		{
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	repoAct := mocks.NewIRepositoryActivities(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindActivities", mock.Anything).Return(&activites, errors.New(""))
	usecaseAct := usecase.NewActivitiesUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindActivities(1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_FindOneActivitiesOk(t *testing.T) {
	activites := entity.Activities{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryActivities(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindActivitiesOne", mock.Anything).Return(&activites, nil)
	usecaseAct := usecase.NewActivitiesUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindActivitiesOne(1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_FindOneActivitiesErr(t *testing.T) {
	activites := entity.Activities{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryActivities(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindActivitiesOne", mock.Anything).Return(&activites, errors.New(""))
	usecaseAct := usecase.NewActivitiesUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindActivitiesOne(1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteActivitiesOk(t *testing.T) {
	activites := entity.Activities{
		Image: "https://adasd.co/asdas",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositoryActivities(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindActivitiesOne", mock.Anything).Return(&activites, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoAct.On("DeleteActivities", mock.Anything).Return(nil)

	usecaseAct := usecase.NewActivitiesUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteActivities(context.Background(), 1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_DeleteActivitiesErrFindOne(t *testing.T) {
	activites := entity.Activities{
		Image: "https://adasd.co/asdas",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositoryActivities(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindActivitiesOne", mock.Anything).Return(&activites, errors.New(""))
	usecaseAct := usecase.NewActivitiesUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteActivities(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteActivitiesErrDeleteFile(t *testing.T) {
	activites := entity.Activities{
		Image: "https://adasd.co/asdas",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositoryActivities(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindActivitiesOne", mock.Anything).Return(&activites, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(""))

	usecaseAct := usecase.NewActivitiesUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteActivities(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteActivitiesErrDeleteAct(t *testing.T) {
	activites := entity.Activities{
		Image: "https://adasd.co/asdas",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositoryActivities(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindActivitiesOne", mock.Anything).Return(&activites, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoAct.On("DeleteActivities", mock.Anything).Return(errors.New(""))

	usecaseAct := usecase.NewActivitiesUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteActivities(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}
