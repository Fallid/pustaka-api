package api

import (
	"errors"
	"fmt"
	"net/http"
	"pustaka-api/exception"
	"pustaka-api/service"
	"pustaka-api/utils"
	"pustaka-api/validator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookHandler struct {
	service service.BookService
}

func NewBookHandler(service service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) GetBooksHandler(ctx *gin.Context) {
	// Get user_id from JWT context
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Status: "error",
			Error:  "User not authenticated",
		})
		return
	}

	result, err := h.service.FindAll(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	booksResponse := utils.ConvertToBooksResponse(result)

	ctx.JSON(http.StatusOK, utils.BooksResponse{
		Status: "success",
		Data:   booksResponse,
	})
}

func (h *BookHandler) GetBookByIdHandler(ctx *gin.Context) {
	bookId := ctx.Param("bookId")

	// Get user_id from JWT context
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Status: "error",
			Error:  "User not authenticated",
		})
		return
	}

	result, err := h.service.FindById(bookId, userID.(string))
	if err != nil {
		// Check if error is record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.Error(exception.NewAppError(http.StatusNotFound, "Book not found or you don't have access"))
			return
		}
		ctx.Error(exception.ErrInternal)
		return
	}

	bookResponse := utils.ConvertToBookResponse(result)

	ctx.JSON(http.StatusOK, utils.BookResponse{
		Status: "success",
		Data:   bookResponse,
	})
}

func (h *BookHandler) PostBookHandler(ctx *gin.Context) {
	var bookInput validator.BookPostRequest

	err := ctx.ShouldBindBodyWithJSON(&bookInput)
	if utils.HandleValidationError(ctx, err) {
		return
	}

	// Get user_id from JWT context
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Status: "error",
			Error:  "User not authenticated",
		})
		return
	}

	createdBook, err := h.service.Create(bookInput, userID.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	bookResponse := utils.ConvertToBookResponse(createdBook)

	ctx.JSON(http.StatusCreated, utils.BookResponse{
		Status: "success",
		Data:   bookResponse,
	})
}

func (h *BookHandler) UpdateBookHandler(ctx *gin.Context) {
	var bookRequest validator.BookUpdateRequest

	bookId := ctx.Param("bookId")

	err := ctx.ShouldBindBodyWithJSON(&bookRequest)
	if utils.HandleValidationError(ctx, err) {
		return
	}

	// Get user_id from JWT context
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Status: "error",
			Error:  "User not authenticated",
		})
		return
	}

	updatedBook, err := h.service.Update(bookId, bookRequest, userID.(string))
	if err != nil {
		// Check if error is authorization error
		if err.Error() == "unauthorized: you are not the owner of this book" {
			ctx.JSON(http.StatusForbidden, utils.ErrorResponse{
				Status: "error",
				Error:  err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	bookResponse := utils.ConvertToBookResponse(updatedBook)

	ctx.JSON(http.StatusOK, utils.BookResponse{
		Status: "success",
		Data:   bookResponse,
	})
}

func (h *BookHandler) DeleteBookHandler(ctx *gin.Context) {
	bookId := ctx.Param("bookId")

	// Get user_id from JWT context
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Status: "error",
			Error:  "User not authenticated",
		})
		return
	}

	err := h.service.Delete(bookId, userID.(string))
	if err != nil {
		// Check if error is authorization error
		if err.Error() == "unauthorized: you are not the owner of this book" {
			ctx.JSON(http.StatusForbidden, utils.ErrorResponse{
				Status: "error",
				Error:  err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.BookDeleteApiResponse{
		Status:  "success",
		Message: fmt.Sprintf("Book dengan id %s, berhasil dihapus", bookId),
	})
}
