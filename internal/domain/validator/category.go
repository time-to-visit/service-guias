package validator

import (
	"net/http"
	"service-user/internal/domain/entity"
	validatorPer "service-user/internal/infra/validation"

	validatorV "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateCategories(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validatorPer.NewValidator()
		categories := new(entity.Categories)

		_ = c.Bind(&categories)
		if err := v.Struct(categories); err != nil {
			errs := err.(validatorV.ValidationErrors)
			return c.JSON(http.StatusBadRequest, validatorPer.GenerateMessage(v, errs))
		}
		c.Set("categories", categories)
		return next(c)
	}
}
