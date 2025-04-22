package models

import (
	"time"

	guuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type BankAccount struct {
	ID           	guuid.UUID       	`gorm:"primaryKey" json:"id"`
	CustomerID   	guuid.UUID       	`json:"customer_id"`
	AccountNumber string						`gorm:"type:varchar(20);unique;not null" json:"account_number"`
	AccountType   string   					`gorm:"type:varchar(50);not null" json:"account_type"` // 'Saving', 'Current', 'Business'
	Balance       float64  					`gorm:"type:decimal(15,2);default:0.00" json:"balance"`
	Currency      string   					`gorm:"type:varchar(10);default:'IDR'" json:"currency"`
	CreatedAt 		time.Time  				`gorm:"autoCreateTime" json:"-"`
  UpdatedAt 		time.Time  				`gorm:"autoUpdateTime" json:"-"`
  DeletedAt 		gorm.DeletedAt 		`gorm:"index" json:"-"` // Menggunakan soft delete
	Customer      *Customer  				`gorm:"foreignKey:CustomerID" json:"customer"`
	TotalBalance  totalBalance			`gorm:"-" json:"totalBalance"`
	Pockets       []Pocket  				`gorm:"foreignKey:BankAccountID" json:"pockets"`
	TermDeposits  []TermDeposit 		`gorm:"foreignKey:BankAccountID" json:"term_deposits"`
}

type BaseBankAccount struct {
	ID           	guuid.UUID       	`gorm:"primaryKey" json:"id"`
	CustomerID   	guuid.UUID       	`json:"customer_id"`
	AccountNumber string						`gorm:"type:varchar(20);unique;not null" json:"account_number"`
	AccountType   string   					`gorm:"type:varchar(50);not null" json:"account_type"` // 'Saving', 'Current', 'Business'
	Balance       float64  					`gorm:"type:decimal(15,2);default:0.00" json:"balance"`
	Currency      string   					`gorm:"type:varchar(10);default:'IDR'" json:"currency"`
	CreatedAt 		time.Time  				`gorm:"autoCreateTime" json:"-"`
  UpdatedAt 		time.Time  				`gorm:"autoUpdateTime" json:"-"`
  DeletedAt 		gorm.DeletedAt 		`gorm:"index" json:"-"` // Menggunakan soft delete
	Pockets       []Pocket  				`gorm:"foreignKey:BankAccountID" json:"pockets"`
	TermDeposits   []TermDeposit    `gorm:"foreignKey:BankAccountID" json:"term_deposits"`
}

type totalBalance struct{
	TotalPocketBalance float64 `gorm:"-" json:"total_pocket_balance"`
	TotalDepositeBalance float64 `gorm:"-" json:"total_deposite_balance"`
	TotalBalance float64 `gorm:"-" json:"total_balance"`
}

type BankAccountResponse struct {
	BankAccount   *BankAccount `json:"bank_account"`
	Customer      *Customer `json:"customer"`
}

type BankAccountList struct {
	TotalCount int     `json:"total_count"`
	TotalPages int     `json:"total_pages"`
	Page       int     `json:"page"`
	Size       int     `json:"size"`
	HasMore    bool    `json:"has_more"`
	BankAccount      []*BankAccount `json:"bank_accounts"`
}

type BankAccountCustomer struct {
	ID           	guuid.UUID       	`gorm:"primaryKey" json:"id"`
	// CustomerID   	guuid.UUID       	`json:"customer_id"`
	AccountNumber string						`gorm:"type:varchar(20);unique;not null" json:"account_number"`
	AccountType   string   					`gorm:"type:varchar(50);not null" json:"account_type"` // 'Saving', 'Current', 'Business'
	Balance       float64  					`gorm:"type:decimal(15,2);default:0.00" json:"balance"`
	Currency      string   					`gorm:"type:varchar(10);default:'IDR'" json:"currency"`
	CreatedAt 		time.Time  				`gorm:"autoCreateTime" json:"-"`
  UpdatedAt 		time.Time  				`gorm:"autoUpdateTime" json:"-"`
  DeletedAt 		gorm.DeletedAt 		`gorm:"index" json:"-"` // Menggunakan soft delete
}

