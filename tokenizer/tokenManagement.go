package tokenizer

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenPayload struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateAccessToken generates a JWT access token with 10 minutes expiration
func GenerateAccessToken(payload TokenPayload, secretKey string) (string, error) {
	claims := Claims{
		UserID:   payload.UserID,
		Username: payload.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GenerateRefreshToken generates a JWT refresh token with 1 week expiration
func GenerateRefreshToken(payload TokenPayload, secretKey string) (string, error) {
	claims := Claims{
		UserID:   payload.UserID,
		Username: payload.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken verifies and parses a JWT token
func VerifyToken(tokenString string, secretKey string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

