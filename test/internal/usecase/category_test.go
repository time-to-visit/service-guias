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

func Test_RegisterCategoryOk(t *testing.T) {
	categories := entity.Categories{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryCategories(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	repoCat.On("InsertCategory", mock.Anything).Return(&categories, nil)
	usecaseAct := usecase.NewCateogriesUseCase(repoCat, objectFile)
	_, status := usecaseAct.RegisterCategory(context.Background(), &categories)
	assert.Equal(t, status, http.StatusOK)
}

func Test_RegisterCategoryErrSetFile(t *testing.T) {
	activites := entity.Categories{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryCategories(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New(""))
	usecaseAct := usecase.NewCateogriesUseCase(repoAct, objectFile)
	_, status := usecaseAct.RegisterCategory(context.Background(), &activites)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_RegisterCategoryErrRegister(t *testing.T) {
	activites := entity.Categories{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryCategories(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	repoAct.On("InsertCategory", mock.Anything).Return(&activites, errors.New(""))
	usecaseAct := usecase.NewCateogriesUseCase(repoAct, objectFile)
	_, status := usecaseAct.RegisterCategory(context.Background(), &activites)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_FindCategoryOk(t *testing.T) {
	activites := []entity.Categories{
		{
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	repoAct := mocks.NewIRepositoryCategories(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindCategory", mock.Anything).Return(&activites, nil)
	usecaseAct := usecase.NewCateogriesUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindCategory(1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_FindCategoryErr(t *testing.T) {
	activites := []entity.Categories{
		{
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	repoAct := mocks.NewIRepositoryCategories(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindCategory", mock.Anything).Return(&activites, errors.New(""))
	usecaseAct := usecase.NewCateogriesUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindCategory(1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_FindOneCategoryOk(t *testing.T) {
	activites := entity.Categories{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryCategories(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindCategoryOne", mock.Anything).Return(&activites, nil)
	usecaseAct := usecase.NewCateogriesUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindCategoryOne(1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_FindOneCategoryErr(t *testing.T) {
	activites := entity.Categories{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryCategories(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindCategoryOne", mock.Anything).Return(&activites, errors.New(""))
	usecaseAct := usecase.NewCateogriesUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindCategoryOne(1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteCategoryOk(t *testing.T) {
	activites := entity.Categories{
		Cover: "https://adasd.co/asdas",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositoryCategories(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindCategoryOne", mock.Anything).Return(&activites, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoAct.On("DeleteCategory", mock.Anything).Return(nil)

	usecaseAct := usecase.NewCateogriesUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteCategory(context.Background(), 1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_DeleteCategoryErrFindOne(t *testing.T) {
	activites := entity.Categories{
		Cover: "https://adasd.co/asdas",

		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositoryCategories(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindCategoryOne", mock.Anything).Return(&activites, errors.New(""))
	usecaseAct := usecase.NewCateogriesUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteCategory(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteCategoryErrDeleteFile(t *testing.T) {
	activites := entity.Categories{
		Cover: "https://adasd.co/asdas",

		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositoryCategories(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindCategoryOne", mock.Anything).Return(&activites, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(""))

	usecaseAct := usecase.NewCateogriesUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteCategory(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteCategoryErrDeleteAct(t *testing.T) {
	activites := entity.Categories{
		Cover: "https://adasd.co/asdas",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositoryCategories(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindCategoryOne", mock.Anything).Return(&activites, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoAct.On("DeleteCategory", mock.Anything).Return(errors.New(""))

	usecaseAct := usecase.NewCateogriesUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteCategory(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}
