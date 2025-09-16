package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/tudemaha/tujuhin-be/pkg/dto/response"
)

func RequestBodyValidator(input interface{}) (response.ArrErrorResponse, bool) {
	validate := validator.New()

	var arrError response.ArrErrorResponse
	if err := validate.Struct(input); err != nil {
		errs := err.(validator.ValidationErrors)

		for _, err := range errs {
			v := response.NewErrorResponseValue(err.Field(), err.Error())
			arrError = append(arrError, v)
		}

		return arrError, true
	}

	return arrError, false
}
