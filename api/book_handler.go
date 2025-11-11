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
	validatorPkg "github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type BookHandler struct {
	service service.BookService
}

func NewBookHandler(service service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) GetBooksHandler(ctx *gin.Context) {
	result, err := h.service.FindAll()
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

	result, err := h.service.FindById(bookId)
	if err != nil {
		// Check if error is record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.Error(exception.NewAppError(http.StatusNotFound, "Book not found"))
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

	createdBook, err := h.service.Create(bookInput)
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
	if err != nil {
		var ve validatorPkg.ValidationErrors
		if errors.As(err, &ve) {
			errorMessages := []string{}
			for _, e := range ve {
				errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
				Status: "error",
				Error:  errorMessages,
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	updatedBook, err := h.service.Update(bookId, bookRequest)
	if err != nil {
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

	err := h.service.Delete(bookId)
	if err != nil {
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
