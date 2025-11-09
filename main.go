package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H {
			"name": "Ahmad Naufal",
			"bio": "Software junior",
		})
	})

	router.Run(":3000")
}