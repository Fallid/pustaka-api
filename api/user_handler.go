package api

import (
	"net/http"
	"pustaka-api/service"
	"pustaka-api/utils"
	"pustaka-api/validator"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
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
