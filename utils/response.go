package utils

import "time"

// Books response
type BookGetData struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Rating      int    `json:"rating"`
	Owner       string `json:"owner"`
}

type BooksResponse struct {
	Status string        `json:"status"`
	Data   []BookGetData `json:"data"`
}

type BookResponse struct {
	Status string      `json:"status"`
	Data   BookGetData `json:"data"`
}

type BookDeleteApiResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Users response
type UserRegisterData struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

type UserRegisterResponse struct {
	Status string           `json:"status"`
	Data   UserRegisterData `json:"data"`
}

type UserGetData struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Fullname  string    `json:"fullname"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserGetResponse struct {
	Status string      `json:"status"`
	Data   UserGetData `json:"data"`
}

// Login response
type UserLoginData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserLoginResponse struct {
	Status string        `json:"status"`
	Data   UserLoginData `json:"data"`
}

// Refresh token response
type RefreshTokenData struct {
	AccessToken string `json:"access_token"`
}

type RefreshTokenResponse struct {
	Status string           `json:"status"`
	Data   RefreshTokenData `json:"data"`
}

// Logout response
type LogoutResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status string      `json:"status"`
	Error  interface{} `json:"error"`
}
