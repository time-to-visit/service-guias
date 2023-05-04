package entry

import (
	"net/http"
	"service-user/internal/domain/entity"
	"service-user/internal/domain/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type questionEntry struct {
	questionCaseuse usecase.QuestionUseCase
}

func NewQuestionEntry(questionCaseuse usecase.QuestionUseCase) *questionEntry {
	return &questionEntry{
		questionCaseuse,
	}
}

func (c *questionEntry) InsertQuestion(context echo.Context) error {
	question := context.Get("question").(*entity.Question)
	response, status := c.questionCaseuse.InsertQuestion(context.Request().Context(), *question)
	return context.JSON(status, response)
}

func (c *questionEntry) DeleteQuestion(context echo.Context) error {
	id := context.Param("ID")
	idQuestion, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := c.questionCaseuse.DeleteQuestion(context.Request().Context(), int64(idQuestion))
	return context.JSON(status, response)
}

func (c *questionEntry) FindQuestionByObjective(context echo.Context) error {
	id := context.Param("ID")
	idObective, err := strconv.Atoi(id)
	if err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	response, status := c.questionCaseuse.FindQuestionByObjective(int64(idObective))
	return context.JSON(status, response)
}
