package handler

import (
	"service-user/cmd/entry"
	"service-user/internal/domain/usecase"
	"service-user/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerQuestion(e *echo.Echo, questionUseCase usecase.QuestionUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	questionEntry := entry.NewQuestionEntry(questionUseCase)
	e.GET("/question/objective/:ID", questionEntry.FindQuestionByObjective, auth)
	e.POST("/question", questionEntry.InsertQuestion, auth, validator.ValidateQuestion)
	e.DELETE("/question/:ID", questionEntry.DeleteQuestion, auth)
	return e
}
