package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)

	router.GET("/book/:id", booksHandler)

	router.Run(":3000")
}

func rootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Ahmad Naufal",
		"bio":  "Software junior",
	})
}

func booksHandler(ctx *gin.Context)  {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}