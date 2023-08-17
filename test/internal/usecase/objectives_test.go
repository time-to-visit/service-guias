package usecase_test

import (
	"errors"
	"net/http"
	"service-user/internal/domain/entity"
	"service-user/internal/domain/usecase"
	"service-user/mocks"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func Test_RegisterObjectiveOk(t *testing.T) {
	categories := entity.Objectives{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryObjectives(t)
	repoCat.On("InsertObjective", mock.Anything).Return(&categories, nil)
	usecaseAct := usecase.NewObjectivesUseCase(repoCat)
	_, status := usecaseAct.RegisterObjective(&categories)
	assert.Equal(t, status, http.StatusOK)
}

func Test_RegisterObjectivesErrRegister(t *testing.T) {
	activites := entity.Objectives{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryObjectives(t)
	repoAct.On("InsertObjective", mock.Anything).Return(&activites, errors.New(""))
	usecaseAct := usecase.NewObjectivesUseCase(repoAct)
	_, status := usecaseAct.RegisterObjective(&activites)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_FindObjectivesOk(t *testing.T) {
	activites := []entity.Objectives{
		{
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	repoAct := mocks.NewIRepositoryObjectives(t)
	repoAct.On("FindObjectiveByLevel", mock.Anything).Return(&activites, nil)
	usecaseAct := usecase.NewObjectivesUseCase(repoAct)
	_, status := usecaseAct.FindObjectiveByLevels(1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_FindObjectivesErr(t *testing.T) {
	activites := []entity.Objectives{
		{
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	repoAct := mocks.NewIRepositoryObjectives(t)
	repoAct.On("FindObjectiveByLevel", mock.Anything).Return(&activites, errors.New(""))
	usecaseAct := usecase.NewObjectivesUseCase(repoAct)
	_, status := usecaseAct.FindObjectiveByLevels(1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_FindOneObjectivesOk(t *testing.T) {
	activites := entity.Objectives{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryObjectives(t)
	repoAct.On("FindObjectiveOne", mock.Anything).Return(&activites, nil)
	usecaseAct := usecase.NewObjectivesUseCase(repoAct)
	_, status := usecaseAct.FindObjectiveOne(1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_FindOneObjectivesErr(t *testing.T) {
	activites := entity.Objectives{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryObjectives(t)
	repoAct.On("FindObjectiveOne", mock.Anything).Return(&activites, errors.New(""))
	usecaseAct := usecase.NewObjectivesUseCase(repoAct)
	_, status := usecaseAct.FindObjectiveOne(1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteObjectivesOk(t *testing.T) {

	repoAct := mocks.NewIRepositoryObjectives(t)
	repoAct.On("DeleteObjective", mock.Anything).Return(nil)

	usecaseAct := usecase.NewObjectivesUseCase(repoAct)
	_, status := usecaseAct.DeleteObjective(1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_DeleteObjectivesErrDeleteAct(t *testing.T) {

	repoAct := mocks.NewIRepositoryObjectives(t)
	repoAct.On("DeleteObjective", mock.Anything).Return(errors.New(""))

	usecaseAct := usecase.NewObjectivesUseCase(repoAct)
	_, status := usecaseAct.DeleteObjective(1)
	assert.Equal(t, status, http.StatusBadRequest)
}
