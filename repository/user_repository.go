package repository

import (
	"pustaka-api/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User)(models.User, error)
	FindById(id string)(models.User, error)
	FindByUsername(username string)(models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return  &userRepository{db: db}
}

func (r *userRepository) Create(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) FindById(id string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "id = ?", id).Error
	return user, err
}

func (r *userRepository) FindByUsername(username string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "username = ?", username).Error
	return user, err
}
