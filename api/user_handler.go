package api

import (
	"errors"
	"net/http"
	"pustaka-api/config"
	"pustaka-api/exception"
	"pustaka-api/service"
	"pustaka-api/utils"
	"pustaka-api/validator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	service service.UserService
	config  *config.Config
}

func NewUserHandler(service service.UserService, cfg *config.Config) *UserHandler {
	return &UserHandler{
		service: service,
		config:  cfg,
	}
}

func (h *UserHandler) RegisterHandler(ctx *gin.Context) {
	var request validator.UserPostRequest

	err := ctx.ShouldBindJSON(&request)
	if utils.HandleValidationError(ctx, err) {
		return
	}

	createUser, err := h.service.Create(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	response := utils.ConvertToRegisterResponse(createUser)

	ctx.JSON(http.StatusCreated, utils.UserRegisterResponse{
		Status: "success",
		Data:   response,
	})
}

func (h *UserHandler) GetUserByIdHandler(ctx *gin.Context) {
	userId := ctx.Param("userId")

	result, err := h.service.FindById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.Error(exception.NewAppError(http.StatusNotFound, "User not found"))
			return
		}
		ctx.Error(exception.ErrInternal)
		return
	}

	userResponse := utils.ConvertToUserGetResponse(result)

	ctx.JSON(http.StatusOK, utils.UserGetResponse{
		Status: "success",
		Data: userResponse,
	})
}

func (h *UserHandler) LoginHandler(ctx *gin.Context) {
	var request validator.UserLoginRequest

	err := ctx.ShouldBindJSON(&request)
	if utils.HandleValidationError(ctx, err) {
		return
	}

	accessToken, refreshToken, err := h.service.Login(request, h.config.AccessTokenKey, h.config.RefreshTokenKey)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	response := utils.ConvertToLoginResponse(accessToken, refreshToken)

	ctx.JSON(http.StatusOK, utils.UserLoginResponse{
		Status: "success",
		Data:   response,
	})
}

func (h *UserHandler) RefreshTokenHandler(ctx *gin.Context) {
	var request validator.RefreshTokenRequest

	err := ctx.ShouldBindJSON(&request)
	if utils.HandleValidationError(ctx, err) {
		return
	}

	newAccessToken, err := h.service.RefreshToken(request.RefreshToken, h.config.AccessTokenKey, h.config.RefreshTokenKey)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.RefreshTokenResponse{
		Status: "success",
		Data: utils.RefreshTokenData{
			AccessToken: newAccessToken,
		},
	})
}

func (h *UserHandler) LogoutHandler(ctx *gin.Context) {
	var request validator.RefreshTokenRequest

	err := ctx.ShouldBindJSON(&request)
	if utils.HandleValidationError(ctx, err) {
		return
	}

	err = h.service.Logout(request.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.LogoutResponse{
		Status:  "success",
		Message: "Successfully logged out",
	})
}
