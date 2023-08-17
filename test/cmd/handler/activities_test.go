package handler_test

import (
	"service-user/cmd/handler"
	"service-user/internal/domain/usecase"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_HandlerActivities(t *testing.T) {
	e := handler.NewHandlerActivities(echo.New(), usecase.ActivitiesUseCase{}, func(next echo.HandlerFunc) echo.HandlerFunc { return nil })
	assert.NotNil(t, e)
}
