package dto

import (

	// validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateBARequest struct {
	AccountNumber string						`json:"account_number"`
	AccountType   string   					`json:"account_type"` // 'Saving', 'Current', 'Business'
	Balance       float64  					`json:"balance"`
	Currency      string   					`json:"currency"`
}

type UpdateBARequest struct {
	AccountNumber string						`json:"account_number"`
	AccountType   string   					`json:"account_type"` // 'Saving', 'Current', 'Business'
	Balance       float64  					`json:"balance"`
	Currency      string   					`json:"currency"`
}