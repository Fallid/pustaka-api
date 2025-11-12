package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID string `gorm:"primaryKey;type:varchar(50)" json:"id"`
	Username string `gorm:"unique;not null;type:varchar(100)" json:"username"`
	Password string `gorm:"not null;type:varchar(255)" json:"password"`
	Fullname string `gorm:"not null;type:varchar(255)" json:"fullname"`
	Books []Book `gorm:"foreignKey:OwnerID" json:"books"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
