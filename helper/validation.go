package helper

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

type ValidationError struct {
	InnerError ValidationErrorData

	Err error
}

func (r *ValidationError) Error() string {
	return r.Err.Error()
}

type ValidationErrorData map[string][]string

func ValidateStruct(data interface{}) error {
	err := validator.New().Struct(data)

	if err != nil {
		valErr := ValidationErrorData{}

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

		return &ValidationError{
			InnerError: valErr,
			Err:        fmt.Errorf("validation faild %w", err),
		}
	}

	return nil
}
