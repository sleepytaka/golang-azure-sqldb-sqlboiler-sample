package validations

import (
	"github.com/go-playground/validator/v10"
	"strconv"
)

func MoneyValidation(fl validator.FieldLevel) bool {
	_, err := strconv.ParseFloat(fl.Field().String(), 64)
	if err != nil {
		return false
	}
	return true
}
