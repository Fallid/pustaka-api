package utils

import "pustaka-api/models"

// ConvertToBookResponse converts Book model to BookResponse
func ConvertToBookResponse(book models.Book) BookGetData {
	return BookGetData{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Rating:      book.Rating,
		Price:       book.Price,
	}
}

// ConvertToBooksResponse converts slice of Book models to slice of BookResponse
func ConvertToBooksResponse(books []models.Book) []BookGetData {
	var booksResponse []BookGetData
	for _, book := range books {
		booksResponse = append(booksResponse, ConvertToBookResponse(book))
	}
	return booksResponse
}

// ConvertToRegisterResponse convert User model to RegisterResponse
func ConvertToRegisterResponse(user models.User) UserRegisterData {
	return  UserRegisterData{
		UserID: user.ID,
		Username: user.Username,
	}
}