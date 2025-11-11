package exception

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// Common errors
var (
	ErrNotFound     = NewAppError(http.StatusNotFound, "Resource not found")
	ErrBadRequest   = NewAppError(http.StatusBadRequest, "Bad request")
	ErrUnauthorized = NewAppError(http.StatusUnauthorized, "Unauthorized")
	ErrInternal     = NewAppError(http.StatusInternalServerError, "Internal server error")
)

// ErrorHandler middleware untuk handle error secara global
func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		// Check if there are any errors
		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last()

			// Check if it's our custom error
			if appErr, ok := err.Err.(*AppError); ok {
				ctx.JSON(appErr.Code, gin.H{
					"status": "error",
					"error":  appErr.Message,
				})
				return
			}

			// Default error response
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
		}
	}
}
