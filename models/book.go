package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          string         `gorm:"primaryKey;type:varchar(50)" json:"id"`
	Title       string         `gorm:"type:varchar(255)" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	Price       int64          `gorm:"type:bigint" json:"price"`
	Rating      int            `gorm:"type:int" json:"rating"`
	OwnerID     string         `gorm:"column:user_id;type:varchar(50);not null" json:"owner_id"`
	User        User           `gorm:"foreignKey:OwnerID;references:ID" json:"user"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
