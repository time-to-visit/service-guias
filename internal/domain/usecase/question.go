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

type QuestionUseCase struct {
	repoQuestion repository.IRepositoryQuestion
	file         storage.IGCImageRepo
}

func NewQuestionUseCase(repoQuestion repository.IRepositoryQuestion, file storage.IGCImageRepo) QuestionUseCase {
	return QuestionUseCase{
		repoQuestion,
		file,
	}
}

func (c *QuestionUseCase) InsertQuestion(ctx context.Context, question entity.Question) (interface{}, int) {
	if question.Type == "IMAGE" {
		pathname, err := c.file.SetFile(ctx, question.Request, "question/question-%s.png")
		if err != nil {
			return objectValues.NewResponseWithData(http.StatusBadRequest,
				"error",
				"hubo un problema con storage",
				nil,
			), http.StatusBadRequest
		}
		question.Request = pathname
	}
	newQuestion, err := c.repoQuestion.InsertQuestion(question)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al insertar new question", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "question insertado exitosamente", newQuestion), http.StatusOK

}

func (c *QuestionUseCase) DeleteQuestion(ctx context.Context, idQuestion int64) (interface{}, int) {
	questions, err := c.repoQuestion.FindQuestionByObjective(idQuestion)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}

	for _, question := range *questions {
		if question.Type == "IMAGE" {
			objectName := utils.ExtractObjectName(question.Request)
			err = c.file.DeleteFile(ctx, "question/%s", objectName)
			if err != nil {
				return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar el contenido del nivel", nil), http.StatusBadRequest
			}
		}
	}

	err = c.repoQuestion.DeleteQuestion(idQuestion)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar new question", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "question eliminada exitosamente", nil), http.StatusOK
}

func (c *QuestionUseCase) FindQuestionByObjective(idObjective int64) (interface{}, int) {
	question, err := c.repoQuestion.FindQuestionByObjective(idObjective)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", question), http.StatusOK
}
