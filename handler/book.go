package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"pustaka-api/books"
)

func RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Ahmad Naufal",
		"bio":  "Software junior",
	})
}

func BooksHandler(ctx *gin.Context) {
	id := ctx.Param("bookId")
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func AuthorsHandler(ctx *gin.Context) {
	bookId := ctx.Param("bookId")
	authorId := ctx.Param("authorId")

	ctx.JSON(http.StatusOK, gin.H{
		"authorId": authorId,
		"bookId":   bookId,
	})
}

func QueryHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	author := ctx.Query("author")
	ctx.JSON(http.StatusOK, gin.H{"title": title, "author": author})
}

func PostBookHandler(ctx *gin.Context) {
	var bookInput books.BookInput

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