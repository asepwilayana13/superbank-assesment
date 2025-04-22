package models

import (
	"time"

	guuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        	guuid.UUID 			`gorm:"primaryKey" json:"-"`
	Username  	string     			`json:"username"`
	Email     	string     			`json:"email"`
	Password  	string     			`json:"-"`
	Role				string					`json:"role"`
	Sessions  	[]Session  			`gorm:"foreignKey:UserRefer; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
	Customer  	[]Customer 			`gorm:"foreignKey:UserID" json:"customer"`
	CreatedAt 	time.Time  			`gorm:"autoCreateTime" json:"-"`
  UpdatedAt 	time.Time  			`gorm:"autoUpdateTime" json:"-"`
  DeletedAt 	gorm.DeletedAt 	`gorm:"index" json:"-"` // Menggunakan soft delete
}

type UsersList struct {
	TotalCount int     `json:"total_count"`
	TotalPages int     `json:"total_pages"`
	Page       int     `json:"page"`
	Size       int     `json:"size"`
	HasMore    bool    `json:"has_more"`
	Users      []*User `json:"users"`
}
