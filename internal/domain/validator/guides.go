package validator

import (
	"net/http"
	"service-user/internal/domain/entity"
	validatorPer "service-user/internal/infra/validation"

	validatorV "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateGuides(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validatorPer.NewValidator()
		guides := new(entity.Guides)

		_ = c.Bind(&guides)
		if err := v.Struct(guides); err != nil {
			errs := err.(validatorV.ValidationErrors)
			return c.JSON(http.StatusBadRequest, validatorPer.GenerateMessage(v, errs))
		}
		c.Set("guides", guides)
		return next(c)
	}
}
