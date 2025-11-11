package service

import (
	"fmt"
	"pustaka-api/models"
	"pustaka-api/repository"
	"pustaka-api/validator"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(request validator.UserPostRequest) (models.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repository: repo}
}

func (s *userService) Create(request validator.UserPostRequest) (models.User, error) {
	userID := fmt.Sprintf("user-%s", uuid.New().String())

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		ID:       userID,
		Username: request.Username,
		Password: string(hash),
		Fullname: request.Fullname,
	}

	newUser, err := s.repository.Create(user)
	if err != nil {
		return models.User{}, err
	}

	return newUser, nil
}