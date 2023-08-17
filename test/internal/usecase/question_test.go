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

func Test_InsertQuestionOk(t *testing.T) {
	categories := entity.Question{
		Type: "IMAGE",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryQuestion(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	repoCat.On("InsertQuestion", mock.Anything).Return(&categories, nil)
	usecaseAct := usecase.NewQuestionUseCase(repoCat, objectFile)
	_, status := usecaseAct.InsertQuestion(context.Background(), categories)
	assert.Equal(t, status, http.StatusOK)
}

func Test_InsertQuestionErrSetFile(t *testing.T) {
	activites := entity.Question{
		Type: "IMAGE",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryQuestion(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New(""))
	usecaseAct := usecase.NewQuestionUseCase(repoAct, objectFile)
	_, status := usecaseAct.InsertQuestion(context.Background(), activites)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_InsertQuestionErrRegister(t *testing.T) {
	activites := entity.Question{
		Type: "IMAGE",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoAct := mocks.NewIRepositoryQuestion(t)
	objectFile := mocks.NewIGCImageRepo(t)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	repoAct.On("InsertQuestion", mock.Anything).Return(&activites, errors.New(""))
	usecaseAct := usecase.NewQuestionUseCase(repoAct, objectFile)
	_, status := usecaseAct.InsertQuestion(context.Background(), activites)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_FindQuestionOk(t *testing.T) {
	activites := []entity.Question{
		{
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	repoAct := mocks.NewIRepositoryQuestion(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindQuestionByObjective", mock.Anything).Return(&activites, nil)
	usecaseAct := usecase.NewQuestionUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindQuestionByObjective(1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_FindQuestionErr(t *testing.T) {
	activites := []entity.Question{
		{
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	repoAct := mocks.NewIRepositoryQuestion(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindQuestionByObjective", mock.Anything).Return(&activites, errors.New(""))
	usecaseAct := usecase.NewQuestionUseCase(repoAct, objectFile)
	_, status := usecaseAct.FindQuestionByObjective(1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteQuestionOk(t *testing.T) {
	activites := []entity.Question{
		{
			Type:    "IMAGE",
			Request: "https://adadasdasd.com/dassa",

			Model: entity.Model{
				ID: 1,
			},
		},
	}

	repoAct := mocks.NewIRepositoryQuestion(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindQuestionByObjective", mock.Anything).Return(&activites, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoAct.On("DeleteQuestion", mock.Anything).Return(nil)

	usecaseAct := usecase.NewQuestionUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteQuestion(context.Background(), 1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_DeleteQuestionErrFindOne(t *testing.T) {
	activites := []entity.Question{
		{
			Type:    "IMAGE",
			Request: "https://adadasdasd.com/dassa",
			Model: entity.Model{
				ID: 1,
			},
		},
	}

	repoAct := mocks.NewIRepositoryQuestion(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindQuestionByObjective", mock.Anything).Return(&activites, errors.New(""))
	usecaseAct := usecase.NewQuestionUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteQuestion(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteQuestionErrDeleteFile(t *testing.T) {
	activites := []entity.Question{
		{
			Type:    "IMAGE",
			Request: "https://adadasdasd.com/dassa",
			Model: entity.Model{
				ID: 1,
			},
		},
	}

	repoAct := mocks.NewIRepositoryQuestion(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindQuestionByObjective", mock.Anything).Return(&activites, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(""))

	usecaseAct := usecase.NewQuestionUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteQuestion(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteQuestionErrDeleteAct(t *testing.T) {
	activites := []entity.Question{
		{
			Type:    "IMAGE",
			Request: "https://adadasdasd.com/dassa",
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	repoAct := mocks.NewIRepositoryQuestion(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoAct.On("FindQuestionByObjective", mock.Anything).Return(&activites, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoAct.On("DeleteQuestion", mock.Anything).Return(errors.New(""))

	usecaseAct := usecase.NewQuestionUseCase(repoAct, objectFile)
	_, status := usecaseAct.DeleteQuestion(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}
