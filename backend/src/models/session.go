package models

import (
	"time"

	guuid "github.com/google/uuid"
)

type Session struct {
	Sessionid guuid.UUID `gorm:"primaryKey" json:"sessionid"`
	Expires   time.Time  `json:"expires"`
	UserRefer guuid.UUID `json:"-"`
	Token string `json:"token"`
	CreatedAt int64      `gorm:"autoCreateTime" json:"-" `
}
