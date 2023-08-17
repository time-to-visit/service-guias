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

func Test_InsertSectionOk(t *testing.T) {
	categories := entity.Section{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositorySection(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	repoCat.On("InsertSection", mock.Anything).Return(&categories, nil)
	usecaseAct := usecase.NewSectionUseCase(repoCat, objectFile)
	_, status := usecaseAct.InsertSection(context.Background(), &categories)
	assert.Equal(t, status, http.StatusOK)
}

func Test_InsertSectionErrSetFile(t *testing.T) {
	activites := entity.Section{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositorySection(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New(""))
	usecaseAct := usecase.NewSectionUseCase(repoAct, objectFile)
	_, status := usecaseAct.InsertSection(context.Background(), &activites)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_InsertSectionErrRegister(t *testing.T) {
	activites := entity.Section{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositorySection(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	repoAct.On("InsertSection", mock.Anything).Return(&activites, errors.New(""))
	usecaseAct := usecase.NewSectionUseCase(repoAct, objectFile)
	_, status := usecaseAct.InsertSection(context.Background(), &activites)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_FindSectionOk(t *testing.T) {
	activites := []entity.Section{
		{
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	repoAct := mocks.NewIRepositorySection(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindSection", mock.Anything).Return(&activites, nil)
	usecaseAct := usecase.NewSectionUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindSection(map[string]interface{}{})
	assert.Equal(t, status, http.StatusOK)
}

func Test_FindSectionErr(t *testing.T) {
	activites := []entity.Section{
		{
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	repoAct := mocks.NewIRepositorySection(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindSection", mock.Anything).Return(&activites, errors.New(""))
	usecaseAct := usecase.NewSectionUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindSection(map[string]interface{}{})
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_FindOneSectionOk(t *testing.T) {
	activites := entity.Section{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositorySection(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindSectionOne", mock.Anything).Return(&activites, nil)
	usecaseAct := usecase.NewSectionUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindSectionOne(1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_FindOneSectionErr(t *testing.T) {
	activites := entity.Section{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositorySection(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindSectionOne", mock.Anything).Return(&activites, errors.New(""))
	usecaseAct := usecase.NewSectionUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindSectionOne(1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteSectionOk(t *testing.T) {
	activites := entity.Section{
		Cover: "https://adasd.co/asdas",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositorySection(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindSectionOne", mock.Anything).Return(&activites, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoAct.On("DeleteSection", mock.Anything).Return(nil)

	usecaseAct := usecase.NewSectionUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteSection(context.Background(), 1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_DeleteSectionErrFindOne(t *testing.T) {
	activites := entity.Section{
		Cover: "https://adasd.co/asdas",

		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositorySection(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindSectionOne", mock.Anything).Return(&activites, errors.New(""))
	usecaseAct := usecase.NewSectionUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteSection(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteSectionErrDeleteFile(t *testing.T) {
	activites := entity.Section{
		Cover: "https://adasd.co/asdas",

		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositorySection(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindSectionOne", mock.Anything).Return(&activites, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(""))

	usecaseAct := usecase.NewSectionUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteSection(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteSectionErrDeleteAct(t *testing.T) {
	activites := entity.Section{
		Cover: "https://adasd.co/asdas",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoAct := mocks.NewIRepositorySection(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindSectionOne", mock.Anything).Return(&activites, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoAct.On("DeleteSection", mock.Anything).Return(errors.New(""))

	usecaseAct := usecase.NewSectionUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteSection(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}
