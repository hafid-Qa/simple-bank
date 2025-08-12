package api

import (
	"go_api/utils"

	"github.com/go-playground/validator/v10"
)

var validCurrencies validator.Func = func(fl validator.FieldLevel) bool {
	currency := fl.Field().String()
	return utils.IsSupportedCurrency(currency)
}
