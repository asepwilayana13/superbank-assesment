package utils

import (
	"github.com/go-ozzo/ozzo-validation/v4"
)

// Massage error Validation
func ValidationErrors(err error) map[string]string {
	errors := make(map[string]string)

	if errs, ok := err.(validation.Errors); ok {
		for field, e := range errs {
			errors[field] = e.Error()
		}
	}

	return errors
}
