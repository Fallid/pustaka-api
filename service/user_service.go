package service

import (
	"errors"
	"fmt"
	"pustaka-api/models"
	"pustaka-api/repository"
	"pustaka-api/tokenizer"
	"pustaka-api/validator"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(request validator.UserPostRequest) (models.User, error)
	FindById(id string) (models.User, error)
	FindByUsername(username string) (models.User, error)
	Login(request validator.UserLoginRequest, accessTokenKey, refreshTokenKey string) (string, string, error)
	RefreshToken(refreshToken string, accessTokenKey, refreshTokenKey string) (string, error)
	Logout(refreshToken string) error
}

type userService struct {
	repository                repository.UserRepository
	authenticationRepository repository.AuthenticationRepository
}

func NewUserService(repo repository.UserRepository, authRepo repository.AuthenticationRepository) UserService {
	return &userService{
		repository:                repo,
		authenticationRepository: authRepo,
	}
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

func (s *userService) FindById(id string) (models.User, error) {
	user, err := s.repository.FindById(id)
	return user, err
}

func (s *userService) FindByUsername(username string) (models.User, error) {
	user, err := s.repository.FindByUsername(username)
	return user, err
}

func (s *userService) Login(request validator.UserLoginRequest, accessTokenKey, refreshTokenKey string) (string, string, error) {
	// Find user by username
	user, err := s.repository.FindByUsername(request.Username)
	if err != nil {
		return "", "", errors.New("invalid username or password")
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return "", "", errors.New("invalid username or password")
	}

	// Create token payload
	payload := tokenizer.TokenPayload{
		UserID:   user.ID,
		Username: user.Username,
	}

	// Generate access token
	accessToken, err := tokenizer.GenerateAccessToken(payload, accessTokenKey)
	if err != nil {
		return "", "", err
	}

	// Generate refresh token
	refreshToken, err := tokenizer.GenerateRefreshToken(payload, refreshTokenKey)
	if err != nil {
		return "", "", err
	}

	// Delete old refresh token if exists
	s.authenticationRepository.DeleteByUserID(user.ID)

	// Save refresh token to database
	refreshTokenModel := models.Authentication{
		ID:        fmt.Sprintf("token-%s", uuid.New().String()),
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	_, err = s.authenticationRepository.Create(refreshTokenModel)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *userService) RefreshToken(refreshToken string, accessTokenKey, refreshTokenKey string) (string, error) {
	// Verify refresh token
	claims, err := tokenizer.VerifyToken(refreshToken, refreshTokenKey)
	if err != nil {
		return "", errors.New("invalid or expired refresh token")
	}

	// Check if refresh token exists in database
	storedToken, err := s.authenticationRepository.FindByToken(refreshToken)
	if err != nil {
		return "", errors.New("refresh token not found")
	}

	// Check if token is expired
	if time.Now().After(storedToken.ExpiresAt) {
		// Delete expired token
		s.authenticationRepository.DeleteByUserID(storedToken.UserID)
		return "", errors.New("refresh token expired")
	}

	// Generate new access token
	payload := tokenizer.TokenPayload{
		UserID:   claims.UserID,
		Username: claims.Username,
	}

	newAccessToken, err := tokenizer.GenerateAccessToken(payload, accessTokenKey)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}

func (s *userService) Logout(refreshToken string) error {
	// Find refresh token in database
	storedToken, err := s.authenticationRepository.FindByToken(refreshToken)
	if err != nil {
		return errors.New("refresh token not found")
	}

	// Delete refresh token from database
	err = s.authenticationRepository.DeleteByUserID(storedToken.UserID)
	if err != nil {
		return err
	}

	return nil
}
