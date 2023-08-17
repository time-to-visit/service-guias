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

func Test_InsertLevelsOk(t *testing.T) {
	categories := entity.Levels{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	repoCat.On("InsertLevels", mock.Anything).Return(&categories, nil)
	usecaseAct := usecase.NewLevelsUseCase(repoCat, objectFile)
	_, status := usecaseAct.InsertLevels(context.Background(), categories)
	assert.Equal(t, status, http.StatusOK)
}

func Test_InsertLevelsErrSetFile(t *testing.T) {
	categories := entity.Levels{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New(""))
	usecaseAct := usecase.NewLevelsUseCase(repoAct, objectFile)
	_, status := usecaseAct.InsertLevels(context.Background(), categories)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_InsertLevelsErrRegister(t *testing.T) {
	categories := entity.Levels{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	repoAct.On("InsertLevels", mock.Anything).Return(&categories, errors.New(""))
	usecaseAct := usecase.NewLevelsUseCase(repoAct, objectFile)
	_, status := usecaseAct.InsertLevels(context.Background(), categories)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_FindLevelsOk(t *testing.T) {
	categories := []entity.Levels{
		{
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	repoAct := mocks.NewIRepositoryLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindLevelsByCategory", mock.Anything).Return(&categories, nil)
	usecaseAct := usecase.NewLevelsUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindLevelsByCategory(1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_FindLevelsErr(t *testing.T) {
	categories := []entity.Levels{
		{
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	repoAct := mocks.NewIRepositoryLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindLevelsByCategory", mock.Anything).Return(&categories, errors.New(""))
	usecaseAct := usecase.NewLevelsUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindLevelsByCategory(1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_FindOneLevelsOk(t *testing.T) {
	categories := entity.Levels{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindLevelsOne", mock.Anything).Return(&categories, nil)
	usecaseAct := usecase.NewLevelsUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindLevelsOne(1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_FindOneLevelsErr(t *testing.T) {
	categories := entity.Levels{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindLevelsOne", mock.Anything).Return(&categories, errors.New(""))
	usecaseAct := usecase.NewLevelsUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindLevelsOne(1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteLevelsOk(t *testing.T) {
	categories := entity.Levels{
		Cover: "http://adasdasd.com/dsadsa",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositoryLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindLevelsOne", mock.Anything).Return(&categories, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoAct.On("DeleteLevels", mock.Anything).Return(nil)

	usecaseAct := usecase.NewLevelsUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteLevels(context.Background(), 1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_DeleteLevelsErrFindOne(t *testing.T) {
	categories := entity.Levels{
		Cover: "http://adasdasd.com/dsadsa",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindLevelsOne", mock.Anything).Return(&categories, errors.New(""))
	usecaseAct := usecase.NewLevelsUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteLevels(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteLevelsErrDeleteFile(t *testing.T) {
	categories := entity.Levels{
		Cover: "http://adasdasd.com/dsadsa",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositoryLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindLevelsOne", mock.Anything).Return(&categories, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(""))

	usecaseAct := usecase.NewLevelsUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteLevels(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteLevelsErrDeleteAct(t *testing.T) {
	categories := entity.Levels{
		Cover: "http://adasdasd.com/dsadsa",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositoryLevels(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindLevelsOne", mock.Anything).Return(&categories, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoAct.On("DeleteLevels", mock.Anything).Return(errors.New(""))

	usecaseAct := usecase.NewLevelsUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteLevels(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}
