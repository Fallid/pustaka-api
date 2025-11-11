package service

import (
	"fmt"
	"pustaka-api/models"
	"pustaka-api/repository"
	"pustaka-api/validator"

	"github.com/google/uuid"
)

type BookService interface {
	FindAll() ([]models.Book, error)
	FindById(id string) (models.Book, error)
	Create(request validator.BookPostRequest) (models.Book, error)
	Update(id string, request validator.BookUpdateRequest) (models.Book, error)
	Delete(id string) error
}

type bookService struct {
	repository repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repository: repo}
}

func (s *bookService) FindAll() ([]models.Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *bookService) FindById(id string) (models.Book, error) {
	book, err := s.repository.FindById(id)
	return book, err
}

func (s *bookService) Create(request validator.BookPostRequest) (models.Book, error) {
	// Generate ID dengan format book-uuid
	bookID := fmt.Sprintf("book-%s", uuid.New().String())

	book := models.Book{
		ID:          bookID,
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		Rating:      request.Rating,
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *bookService) Update(id string, request validator.BookUpdateRequest) (models.Book, error) {
	book, err := s.repository.FindById(id)
	if err != nil {
		return book, err
	}

	book.Title = request.Title
	book.Description = request.Description
	book.Price = request.Price
	book.Rating = request.Rating

	updatedBook, err := s.repository.Update(book)
	return updatedBook, err
}

func (s *bookService) Delete(id string) error {
	err := s.repository.Delete(id)
	return err
}
