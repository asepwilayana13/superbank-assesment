package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateLoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

// Validate method untuk validasi input
func (r CreateLoginRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}
