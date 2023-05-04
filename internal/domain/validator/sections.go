package validator

import (
	"net/http"
	"service-user/internal/domain/entity"
	validatorPer "service-user/internal/infra/validation"

	validatorV "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateSections(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validatorPer.NewValidator()
		section := new(entity.Section)

		_ = c.Bind(&section)
		if err := v.Struct(section); err != nil {
			errs := err.(validatorV.ValidationErrors)
			return c.JSON(http.StatusBadRequest, validatorPer.GenerateMessage(v, errs))
		}
		c.Set("section", section)
		return next(c)
	}
}
