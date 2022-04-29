package helper

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

type ValidationError map[string][]string

func ValidateStruct(data interface{}) *ValidationError {
	err := validator.New().Struct(data)

	if err != nil {
		valErr := ValidationError{}

		//combine all errors
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())

			var msg string

			switch err.Tag() {
			case "required":
				msg = fmt.Sprintf("%s is required", field)
			default:
				msg = fmt.Sprintf("%s failed on %s validation", field, err.Tag())
			}

			valErr[field] = append(valErr[field], msg)
		}

		return &valErr
	}

	return nil
}
