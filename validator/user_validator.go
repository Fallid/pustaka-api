package validator

type UserPostRequest struct {
	Username string `json:"username" binding:"required,max=100"`
	Password string `json:"password" binding:"required,min=6,max=255"`
	Fullname string `json:"fullname" binding:"required,max=255"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required,max=100"`
	Password string `json:"password" binding:"required,min=6,max=255"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}