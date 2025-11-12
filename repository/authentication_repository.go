package repository

import (
	"pustaka-api/models"

	"gorm.io/gorm"
)

type AuthenticationRepository interface {
	Create(token models.Authentication) (models.Authentication, error)
	FindByToken(token string) (models.Authentication, error)
	FindByUserID(userID string) (models.Authentication, error)
	DeleteByUserID(userID string) error
}

type authenticationRepository struct {
	db *gorm.DB
}

func NewAuthenticationRepository(db *gorm.DB) AuthenticationRepository {
	return &authenticationRepository{db: db}
}

func (r *authenticationRepository) Create(token models.Authentication) (models.Authentication, error) {
	err := r.db.Create(&token).Error
	return token, err
}

func (r *authenticationRepository) FindByToken(token string) (models.Authentication, error) {
	var authentication models.Authentication
	err := r.db.First(&authentication, "token = ?", token).Error
	return authentication, err
}

func (r *authenticationRepository) FindByUserID(userID string) (models.Authentication, error) {
	var authentication models.Authentication
	err := r.db.First(&authentication, "user_id = ?", userID).Error
	return authentication, err
}

func (r *authenticationRepository) DeleteByUserID(userID string) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.Authentication{}).Error
}
