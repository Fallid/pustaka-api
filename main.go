package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	// Versioning routes with group
	v1 := router.Group("/v1")
	
	// GET ROUTE
	v1.GET("/", rootHandler)

	v1.GET("/book/:bookId", booksHandler)

	v1.GET("/book/:bookId/author/:authorId", authorsHandler)

	v1.GET("/query", queryHandler)

	// POST ROUTE
	v1.POST("/book", postBookHandler)

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
	Title    string      `json:"title" binding:"required"`
	Price    int `json:"price" binding:"required,number"`
	SubTitle string      `json:"sub_title"` // alias as sub_title
}

func postBookHandler(ctx *gin.Context) {
	var bookInput bookInput

	err := ctx.ShouldBindBodyWithJSON(&bookInput)

	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorMessages := []string{}
			for _, e := range ve {
				errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": "fail",
				"errors": errorMessages,
			})
			return
		}
		// For other errors (e.g. JSON parse errors)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title":     bookInput.Title,
		"price":     bookInput.Price,
		"sub_title": bookInput.SubTitle,
	})
}
