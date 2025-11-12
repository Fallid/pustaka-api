package middleware

import (
	"net/http"
	"pustaka-api/config"
	"pustaka-api/tokenizer"
	"pustaka-api/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get authorization header
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse{
				Status: "error",
				Error:  "Authorization header is required",
			})
			ctx.Abort()
			return
		}

		// Check if it's a Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse{
				Status: "error",
				Error:  "Invalid authorization header format",
			})
			ctx.Abort()
			return
		}

		token := parts[1]

		// Verify token
		claims, err := tokenizer.VerifyToken(token, cfg.AccessTokenKey)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse{
				Status: "error",
				Error:  "Invalid or expired token",
			})
			ctx.Abort()
			return
		}

		// Set user info to context
		ctx.Set("user_id", claims.UserID)
		ctx.Set("username", claims.Username)

		ctx.Next()
	}
}
