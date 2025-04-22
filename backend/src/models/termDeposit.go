package models

import (
	"time"

	guuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type TermDeposit struct {
	ID            			guuid.UUID      `gorm:"primaryKey" json:"id"`
	BankAccountID 			guuid.UUID      `json:"bank_account_id"`
	DepositAmount 			float64   			`gorm:"type:decimal(15,2);not null" json:"deposit_amount"`
	InterestRate  			float64   			`gorm:"type:decimal(5,2);not null" json:"interest_rate"`
	StartDate     			time.Time 			`json:"start_date"`
	Tenor 							int							`json:"tenor"`
	BalancePlusInterest	float64 				`gorm:"type:decimal(15,2);not null" json:"balance_plus_interest"`
	Status        			string    			`gorm:"type:varchar(50);default:'Active'" json:"status"` // 'Active', 'Closed', 'Matured'
	CreatedAt 					time.Time  			`gorm:"autoCreateTime" json:"-"`
  UpdatedAt 					time.Time  			`gorm:"autoUpdateTime" json:"-"`
  DeletedAt 					gorm.DeletedAt 	`gorm:"index" json:"-"` // Menggunakan soft delete
	BankAccount    BankAccount   `gorm:"foreignKey:BankAccountID;references:ID" json:"-"`
}

func CalculateBalancePlusInterest(depositAmount float64, interestRate float64) float64 {
	return depositAmount * (1 + interestRate/100)
}
