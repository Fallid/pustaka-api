package service

import (
	"fmt"
	"pustaka-api/models"
	"pustaka-api/repository"
	"pustaka-api/validator"

	"github.com/google/uuid"
)

type BookService interface {
	FindAll(ownerID string) ([]models.Book, error)
	FindById(id string, ownerID string) (models.Book, error)
	Create(request validator.BookPostRequest, ownerID string) (models.Book, error)
	Update(id string, request validator.BookUpdateRequest, ownerID string) (models.Book, error)
	Delete(id string, ownerID string) error
}

type bookService struct {
	repository repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repository: repo}
}

func (s *bookService) FindAll(ownerID string) ([]models.Book, error) {
	books, err := s.repository.FindAll(ownerID)
	return books, err
}

func (s *bookService) FindById(id string, ownerID string) (models.Book, error) {
	book, err := s.repository.FindById(id, ownerID)
	return book, err
}

func (s *bookService) Create(request validator.BookPostRequest, ownerID string) (models.Book, error) {
	// Generate ID dengan format book-uuid
	bookID := fmt.Sprintf("book-%s", uuid.New().String())

	book := models.Book{
		ID:          bookID,
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		Rating:      request.Rating,
		OwnerID:     ownerID,
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *bookService) Update(id string, request validator.BookUpdateRequest, ownerID string) (models.Book, error) {
	book, err := s.repository.FindById(id, ownerID)
	if err != nil {
		return book, err
	}

	// Check if user is the owner
	if book.OwnerID != ownerID {
		return book, fmt.Errorf("unauthorized: you are not the owner of this book")
	}

	book.Title = request.Title
	book.Description = request.Description
	book.Price = request.Price
	book.Rating = request.Rating

	updatedBook, err := s.repository.Update(book)
	return updatedBook, err
}

func (s *bookService) Delete(id string, ownerID string) error {
	book, err := s.repository.FindById(id, ownerID)
	if err != nil {
		return err
	}

	// Check if user is the owner
	if book.OwnerID != ownerID {
		return fmt.Errorf("unauthorized: you are not the owner of this book")
	}

	err = s.repository.Delete(id)
	return err
}
