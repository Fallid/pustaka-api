package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	validatorPkg "github.com/go-playground/validator/v10"
)

func HandleValidationError(ctx *gin.Context, err error) bool {
	var ve validatorPkg.ValidationErrors
	if errors.As(err, &ve) {
		errorMessages := []string{}
		for _, e := range ve {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Status: "error",
			Error:  errorMessages,
		})
		return true
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return true
	}
	return false
}
