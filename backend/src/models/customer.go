package models

import (
	"time"

	guuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID          guuid.UUID      `gorm:"primaryKey" json:"id"`
	UserID      guuid.UUID      `json:"user_id"`
	FullName    string    			`gorm:"type:varchar(255);not null" json:"full_name"`
	PhoneNumber string    			`gorm:"type:varchar(20);unique;not null" json:"phone_number"`
	DateOfBirth time.Time 			`gorm:"type:date" json:"date_of_birth"`
	Address     string    			`json:"address"`
	CreatedAt 	time.Time  			`gorm:"autoCreateTime" json:"-"`
  UpdatedAt 	time.Time  			`gorm:"autoUpdateTime" json:"-"`
  DeletedAt 	gorm.DeletedAt 	`gorm:"index" json:"-"` // Menggunakan soft delete
	BankAccount  *BankAccount 	`gorm:"foreignKey:CustomerID" json:"bankAccount"`
}

type CustomersList struct {
	TotalCount int     `json:"total_count"`
	TotalPages int     `json:"total_pages"`
	Page       int     `json:"page"`
	Size       int     `json:"size"`
	HasMore    bool    `json:"has_more"`
	Customer  []*Customer `json:"customer"`
}

type BaseResponse struct {
	Customer []*Customer `json:"customer"`
	User []*User `json:"user"`
}
