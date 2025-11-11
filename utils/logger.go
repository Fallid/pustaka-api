package utils

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger middleware untuk logging setiap request
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Process request
		ctx.Next()

		// Calculate latency
		latency := time.Since(startTime)

		// Get request info
		statusCode := ctx.Writer.Status()
		method := ctx.Request.Method
		path := ctx.Request.URL.Path
		clientIP := ctx.ClientIP()

		// Log request
		log.Printf("[%s] %s %s | Status: %d | Latency: %v | IP: %s",
			method,
			path,
			ctx.Request.Proto,
			statusCode,
			latency,
			clientIP,
		)
	}
}
