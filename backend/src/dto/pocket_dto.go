package dto

import (

	// validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreatePocketRequest struct {
	BankAccountID string       			`json:"bank_account_id"`
	PocketName    string    				`json:"pocket_name"`
	Balance       float64   				`json:"balance"`
	GoalAmount    float64   				`json:"goal_amount"`
}

type UpdatePocketRequest struct {
	BankAccountID string       			`json:"bank_account_id"`
	PocketName    string    				`json:"pocket_name"`
	Balance       float64   				`json:"balance"`
	GoalAmount    float64   				`json:"goal_amount"`
}