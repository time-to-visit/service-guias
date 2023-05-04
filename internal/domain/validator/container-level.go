package validator

import (
	"net/http"
	"service-user/internal/domain/entity"
	validatorPer "service-user/internal/infra/validation"

	validatorV "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateContainerLevels(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validatorPer.NewValidator()
		containerLevels := new(entity.ContainerLevels)

		_ = c.Bind(&containerLevels)
		if err := v.Struct(containerLevels); err != nil {
			errs := err.(validatorV.ValidationErrors)
			return c.JSON(http.StatusBadRequest, validatorPer.GenerateMessage(v, errs))
		}
		c.Set("containerLevels", containerLevels)
		return next(c)
	}
}
