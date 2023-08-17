package handler_test

import (
	"service-user/cmd/handler"
	"service-user/internal/domain/usecase"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_HandlerLevel(t *testing.T) {
	e := handler.NewHandlerLevel(echo.New(), usecase.LevelsUseCase{}, func(next echo.HandlerFunc) echo.HandlerFunc { return nil })
	assert.NotNil(t, e)
}
