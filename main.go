package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)

	router.GET("/book/:bookId", booksHandler)
	
	router.GET("/book/:bookId/author/:authorId", authorsHandler)

	router.Run(":3000")
}

func rootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Ahmad Naufal",
		"bio":  "Software junior",
	})
}

func booksHandler(ctx *gin.Context) {
	id := ctx.Param("bookId")
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func authorsHandler(ctx *gin.Context) {
	bookId := ctx.Param("bookId")
	authorId  := ctx.Param("authorId")

	ctx.JSON(http.StatusOK, gin.H{
		"authorId": authorId,
		"bookId": bookId,
	})
}
