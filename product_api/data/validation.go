package data

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func validateSKU(field validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(field.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}

	return true
}

func (p *Product) Validate() error {
	validator := validator.New()
	validator.RegisterValidation("sku", validateSKU)
	return validator.Struct(p)
}
