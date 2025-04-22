package dto

import (

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateCustomerRequest struct {
	FullName    string    			`json:"full_name"`
	UserName    string    			`json:"username"`
	PhoneNumber string    			`json:"phone_number"`
	DateOfBirth string					`json:"date_of_birth"`
	Address     string    			`json:"address"`
	Email				string					`json:"email"`
  AccountType string					`json:"account_type"`
  Balance			int					`json:"balance"`
  Currency		string					`json:"currency"`
}

// Validasi
func (r CreateCustomerRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.FullName, validation.Required),
		validation.Field(&r.UserName, validation.Required),
		validation.Field(&r.PhoneNumber, validation.Required),
		validation.Field(&r.DateOfBirth, validation.Required), 
		validation.Field(&r.Address, validation.Required),
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.AccountType, validation.Required),
		validation.Field(&r.Balance, validation.Required),
		validation.Field(&r.Currency, validation.Required),
	)
}


type UpdateCustomerRequest struct {
	FullName    string    			`json:"full_name"`
	UserName    string    			`json:"username"`
	PhoneNumber string    			`json:"phone_number"`
	DateOfBirth string					`json:"date_of_birth"`
	Address     string    			`json:"address"`
	Email				string					`json:"email"`
  AccountType string					`json:"account_type"`
  Currency		string					`json:"currency"`
}

// Validasi
func (r UpdateCustomerRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.FullName, validation.Required),
		validation.Field(&r.UserName, validation.Required),
		validation.Field(&r.PhoneNumber, validation.Required),
		validation.Field(&r.DateOfBirth, validation.Required), 
		validation.Field(&r.Address, validation.Required),
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.AccountType, validation.Required),
		validation.Field(&r.Currency, validation.Required),
	)
}