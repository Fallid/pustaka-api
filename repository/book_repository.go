package repository

import (
	"pustaka-api/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll() ([]models.Book, error)
	FindById(id string) (models.Book, error)
	Create(book models.Book) (models.Book, error)
	Update(book models.Book) (models.Book, error)
	Delete(id string) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) FindAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepository) FindById(id string) (models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error
	return book, err
}

func (r *bookRepository) Create(book models.Book) (models.Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *bookRepository) Update(book models.Book) (models.Book, error) {
	err := r.db.Save(&book).Error
	return book, err
}

func (r *bookRepository) Delete(id string) error {
	var book models.Book
	err := r.db.Delete(&book, id).Error
	return err
}
