package validator

import (
	"net/http"
	"service-user/internal/domain/entity"
	validatorPer "service-user/internal/infra/validation"

	validatorV "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateObjectives(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validatorPer.NewValidator()
		objectives := new(entity.Objectives)

		_ = c.Bind(&objectives)
		if err := v.Struct(objectives); err != nil {
			errs := err.(validatorV.ValidationErrors)
			return c.JSON(http.StatusBadRequest, validatorPer.GenerateMessage(v, errs))
		}
		c.Set("objectives", objectives)
		return next(c)
	}
}
