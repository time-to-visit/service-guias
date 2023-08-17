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

func Test_InsertContainerLevelsOk(t *testing.T) {
	categories := entity.ContainerLevels{
		Type: "IMAGE",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryContainerLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	repoCat.On("InsertContainerLevels", mock.Anything).Return(&categories, nil)
	usecaseAct := usecase.NewContainerLevelsUseCase(repoCat, objectFile)
	_, status := usecaseAct.InsertContainerLevels(context.Background(), &categories)
	assert.Equal(t, status, http.StatusOK)
}

func Test_InsertContainerLevelsErrSetFile(t *testing.T) {
	categories := entity.ContainerLevels{
		Type: "IMAGE",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryContainerLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New(""))
	usecaseAct := usecase.NewContainerLevelsUseCase(repoAct, objectFile)
	_, status := usecaseAct.InsertContainerLevels(context.Background(), &categories)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_InsertContainerLevelsErrRegister(t *testing.T) {
	categories := entity.ContainerLevels{
		Type: "IMAGE",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryContainerLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	repoAct.On("InsertContainerLevels", mock.Anything).Return(&categories, errors.New(""))
	usecaseAct := usecase.NewContainerLevelsUseCase(repoAct, objectFile)
	_, status := usecaseAct.InsertContainerLevels(context.Background(), &categories)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_FindContainerLevelsOk(t *testing.T) {
	categories := []entity.ContainerLevels{
		{
			Type: "IMAGE",
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	repoAct := mocks.NewIRepositoryContainerLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindContainerLevels", mock.Anything).Return(&categories, nil)
	usecaseAct := usecase.NewContainerLevelsUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindContainerLevels(1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_FindContainerLevelsErr(t *testing.T) {
	categories := []entity.ContainerLevels{
		{
			Type: "IMAGE",
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	repoAct := mocks.NewIRepositoryContainerLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindContainerLevels", mock.Anything).Return(&categories, errors.New(""))
	usecaseAct := usecase.NewContainerLevelsUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindContainerLevels(1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_FindOneContainerLevelsOk(t *testing.T) {
	categories := entity.ContainerLevels{
		Type: "IMAGE",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryContainerLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindContainerLevelsOne", mock.Anything).Return(&categories, nil)
	usecaseAct := usecase.NewContainerLevelsUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindContainerLevelsOne(1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_FindOneContainerLevelsErr(t *testing.T) {
	categories := entity.ContainerLevels{
		Type: "IMAGE",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryContainerLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindContainerLevelsOne", mock.Anything).Return(&categories, errors.New(""))
	usecaseAct := usecase.NewContainerLevelsUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindContainerLevelsOne(1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteContainerLevelsOk(t *testing.T) {
	categories := entity.ContainerLevels{
		Type:      "IMAGE",
		Container: "http://adasdasd.com/dsadsa",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositoryContainerLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindContainerLevelsOne", mock.Anything).Return(&categories, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoAct.On("DeleteContainerLevels", mock.Anything).Return(nil)

	usecaseAct := usecase.NewContainerLevelsUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteContainerLevels(context.Background(), 1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_DeleteContainerLevelsErrFindOne(t *testing.T) {
	categories := entity.ContainerLevels{
		Type:      "IMAGE",
		Container: "http://adasdasd.com/dsadsa",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryContainerLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindContainerLevelsOne", mock.Anything).Return(&categories, errors.New(""))
	usecaseAct := usecase.NewContainerLevelsUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteContainerLevels(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteContainerLevelsErrDeleteFile(t *testing.T) {
	categories := entity.ContainerLevels{
		Type:      "IMAGE",
		Container: "http://adasdasd.com/dsadsa",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositoryContainerLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindContainerLevelsOne", mock.Anything).Return(&categories, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(""))

	usecaseAct := usecase.NewContainerLevelsUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteContainerLevels(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteContainerLevelsErrDeleteAct(t *testing.T) {
	categories := entity.ContainerLevels{
		Type:      "IMAGE",
		Container: "http://adasdasd.com/dsadsa",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositoryContainerLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindContainerLevelsOne", mock.Anything).Return(&categories, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoAct.On("DeleteContainerLevels", mock.Anything).Return(errors.New(""))

	usecaseAct := usecase.NewContainerLevelsUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteContainerLevels(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}
