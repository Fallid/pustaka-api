package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"pustaka-api/books"
)

type bookHandler struct {
	bookService books.Service
}

func NewBookHandler(bookService books.Service) *bookHandler {
	return &bookHandler{bookService: bookService}
}

func (h *bookHandler) GetBooksHandler(ctx *gin.Context) {
	result, err := h.bookService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	var booksResponse []books.BookResponse

	for _, book := range result {
		bookResponse := response(book)
		booksResponse = append(booksResponse, bookResponse)
	}

	ctx.JSON(http.StatusOK, books.BooksApiResponse{
		Status: "success",
		Data:   booksResponse,
	})
}

func (h *bookHandler) GetBookByIdHandler(ctx *gin.Context) {
	var param = ctx.Param("bookId")
	bookId, _ := strconv.Atoi(param)
	result, err := h.bookService.FindById(bookId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	bookResponse := response(result)

	ctx.JSON(http.StatusOK, books.BookApiResponse{
		Status: "success",
		Data:   bookResponse,
	})
}

func (h *bookHandler) PostBookHandler(ctx *gin.Context) {
	var bookInput books.BookPostRequest

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

	h.bookService.Create(bookInput)

	ctx.JSON(http.StatusOK, gin.H{
		"title":       bookInput.Title,
		"description": bookInput.Description,
		"rating":      bookInput.Rating,
		"price":       bookInput.Price,
	})
}

func (h *bookHandler) UpdateBookHandler(ctx *gin.Context) {
	var bookRequest books.BookUpdateRequest

	var param = ctx.Param("bookId")
	bookId, _ := strconv.Atoi(param)

	err := ctx.ShouldBindBodyWithJSON(&bookRequest)

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

	book, err := h.bookService.Update(bookId, bookRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	ctx.JSON(http.StatusOK, books.BookResponse{
		ID:          bookId,
		Title:       book.Title,
		Description: book.Description,
		Rating:      book.Rating,
		Price:       book.Price,
	})
}

func (h *bookHandler) DeleteBookHandler(ctx *gin.Context) {
	var param = ctx.Param("bookId")
	bookId, _ := strconv.Atoi(param)

	err := h.bookService.Delete(bookId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	ctx.JSON(http.StatusOK, books.BookDeleteApiResponse{
		Status: "success",
		Message: fmt.Sprintf("Book dengan id %d, berhasil dihapus", bookId),
	})

}

func response(book books.Book) books.BookResponse {
	return books.BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Rating:      book.Rating,
		Price:       book.Price,
	}
}
