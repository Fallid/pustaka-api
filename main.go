package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// GET ROUTE
	router.GET("/", rootHandler)

	router.GET("/book/:bookId", booksHandler)

	router.GET("/book/:bookId/author/:authorId", authorsHandler)

	router.GET("/query", queryHandler)

	// POST ROUTE
	router.POST("/book", postBookHandler)

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
	authorId := ctx.Param("authorId")

	ctx.JSON(http.StatusOK, gin.H{
		"authorId": authorId,
		"bookId":   bookId,
	})
}

func queryHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	author := ctx.Query("author")
	ctx.JSON(http.StatusOK, gin.H{"title": title, "author": author})
}

type bookInput struct {
	Title string
	Price int
	SubTitle  string `json:"sub_title"` // alias as sub_title
}

func postBookHandler(ctx *gin.Context) {
	var bookInput bookInput

	err := ctx.ShouldBindBodyWithJSON(&bookInput)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
		"sub_title": bookInput.SubTitle,
	})
}
