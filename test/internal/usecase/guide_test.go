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

func Test_InsertGuidesOk(t *testing.T) {
	categories := entity.Guides{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryGuides(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	repoCat.On("InsertGuides", mock.Anything).Return(&categories, nil)
	usecaseAct := usecase.NewGuidesUseCase(repoCat, objectFile)
	_, status := usecaseAct.RegisterGuides(context.Background(), &categories)
	assert.Equal(t, status, http.StatusOK)
}

func Test_InsertGuidesErrSetFile(t *testing.T) {
	categories := entity.Guides{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryGuides(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New(""))
	usecaseAct := usecase.NewGuidesUseCase(repoAct, objectFile)
	_, status := usecaseAct.RegisterGuides(context.Background(), &categories)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_InsertGuidesErrRegister(t *testing.T) {
	categories := entity.Guides{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryGuides(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	repoAct.On("InsertGuides", mock.Anything).Return(&categories, errors.New(""))
	usecaseAct := usecase.NewGuidesUseCase(repoAct, objectFile)
	_, status := usecaseAct.RegisterGuides(context.Background(), &categories)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_FindGuidesOk(t *testing.T) {
	categories := []entity.Guides{
		{
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	repoAct := mocks.NewIRepositoryGuides(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindGuides", mock.Anything).Return(&categories, nil)
	usecaseAct := usecase.NewGuidesUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindGuides(make(map[string]interface{}))
	assert.Equal(t, status, http.StatusOK)
}

func Test_FindGuidesErr(t *testing.T) {
	categories := []entity.Guides{
		{
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	repoAct := mocks.NewIRepositoryGuides(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindGuides", mock.Anything).Return(&categories, errors.New(""))
	usecaseAct := usecase.NewGuidesUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindGuides(make(map[string]interface{}))
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_FindOneGuidesOk(t *testing.T) {
	categories := entity.Guides{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryGuides(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindGuidesOne", mock.Anything).Return(&categories, nil)
	usecaseAct := usecase.NewGuidesUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindGuidesOne(1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_FindOneGuidesErr(t *testing.T) {
	categories := entity.Guides{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryGuides(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindGuidesOne", mock.Anything).Return(&categories, errors.New(""))
	usecaseAct := usecase.NewGuidesUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindGuidesOne(1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteGuidesOk(t *testing.T) {
	categories := entity.Guides{
		Cover: "http://adasdasd.com/dsadsa",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositoryGuides(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindGuidesOne", mock.Anything).Return(&categories, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoAct.On("DeleteGuides", mock.Anything).Return(nil)

	usecaseAct := usecase.NewGuidesUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteGuides(context.Background(), 1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_DeleteGuidesErrFindOne(t *testing.T) {
	categories := entity.Guides{
		Cover: "http://adasdasd.com/dsadsa",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryGuides(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindGuidesOne", mock.Anything).Return(&categories, errors.New(""))
	usecaseAct := usecase.NewGuidesUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteGuides(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteGuidesErrDeleteFile(t *testing.T) {
	categories := entity.Guides{
		Cover: "http://adasdasd.com/dsadsa",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositoryGuides(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindGuidesOne", mock.Anything).Return(&categories, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(""))

	usecaseAct := usecase.NewGuidesUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteGuides(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteGuidesErrDeleteAct(t *testing.T) {
	categories := entity.Guides{
		Cover: "http://adasdasd.com/dsadsa",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositoryGuides(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindGuidesOne", mock.Anything).Return(&categories, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoAct.On("DeleteGuides", mock.Anything).Return(errors.New(""))

	usecaseAct := usecase.NewGuidesUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteGuides(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}
