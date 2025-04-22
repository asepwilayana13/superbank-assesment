package models

import (
	"time"

	guuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Pocket struct {
	ID            guuid.UUID       	`gorm:"primaryKey" json:"id"`
	BankAccountID guuid.UUID       	`json:"bank_account_id"`
	PocketName    string    				`gorm:"type:varchar(100);not null" json:"pocket_name"`
	Balance       float64   				`gorm:"type:decimal(15,2);default:0.00" json:"balance"`
	GoalAmount    float64   				`gorm:"type:decimal(15,2);default:0.00" json:"goal_amount"`
	CreatedAt 		time.Time  				`gorm:"autoCreateTime" json:"-"`
  UpdatedAt 		time.Time  				`gorm:"autoUpdateTime" json:"-"`
  DeletedAt 		gorm.DeletedAt 		`gorm:"index" json:"-"`
	BankAccount   BankAccount 			`gorm:"foreignKey:BankAccountID" json:"-"`
}

type PocketList struct {
	TotalCount int     `json:"total_count"`
	TotalPages int     `json:"total_pages"`
	Page       int     `json:"page"`
	Size       int     `json:"size"`
	HasMore    bool    `json:"has_more"`
	Pockets      []*Pocket `json:"pockets"`
}
