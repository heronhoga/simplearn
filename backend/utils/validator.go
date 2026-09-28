package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateAndMapRequest(inputStruct any) ([]string, []string, int) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(inputStruct)

	if err != nil {
		var validationErrors validator.ValidationErrors
		errorTags := []string{}
		errorMessages := []string{}
		if errors.As(err, &validationErrors) {

			for _, fieldErr := range validationErrors {
				var errorMessage string
				if fieldErr.Tag() == "required" {
					errorMessage = fieldErr.Field() + " is required"
				} else if fieldErr.Tag() == "email" {
					errorMessage = "invalid email format"
				} else if fieldErr.Tag() == "min" {
					errorMessage = fieldErr.Field() + " required at least " + fieldErr.Param() + " characters"
				} else if fieldErr.Tag() == "max" {
					errorMessage = fieldErr.Field() + " exceeded maximum characters (" + fieldErr.Param() + ")"
				}
				errorTags = append(errorTags, fieldErr.Field())
				errorMessages = append(errorMessages, errorMessage)
			}

			return errorTags, errorMessages, 400
		} else {
			return nil, nil, 500
		}
	}

	return nil, nil, 200
}
