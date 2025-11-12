package models

import (
	"time"

	"gorm.io/gorm"
)

type Authentication struct {
	ID        string         `gorm:"primaryKey;type:varchar(50)" json:"id"`
	UserID    string         `gorm:"not null;type:varchar(50)" json:"user_id"`
	Token     string         `gorm:"not null;type:text;unique" json:"token"`
	ExpiresAt time.Time      `gorm:"not null" json:"expires_at"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
