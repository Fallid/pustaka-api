package main

import (
	"fmt"
	"log"
	"pustaka-api/books"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=100.105.102.88 user=developer password=password dbname=pustaka_api port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}

	fmt.Println("Database connection succeed")

	//migration
	db.AutoMigrate(books.Book{})
	
	//CRUD
	// book := books.Book{}
	// book.Title = "Manusia Lemah"
	// book.Price = 90000
	// book.Rating = 5
	// book.Description = "Lorem ipsum"

	// err = db.Create(&book).Error

	// if err != nil {
	// 	fmt.Print("Error Creating book record")
	// }

	var resultBook books.Book
	err = db.First(&resultBook).Error
	if err != nil {
		fmt.Println("========================")
		fmt.Println("Error finding book record")
		fmt.Println("========================")
	}
	
	fmt.Println("========================")
	fmt.Println("Title: ", resultBook.Title)
	fmt.Println("price: ", resultBook.Price)
	fmt.Println("rating: ", resultBook.Rating)
	fmt.Println("========================")

	var resultBook1 books.Book
	err = db.Last(&resultBook1).Error
	if err != nil {
		fmt.Println("========================")
		fmt.Println("Error finding book record")
		fmt.Println("========================")
	 }
	
	fmt.Println("========================")
	fmt.Println("Title: ", resultBook1.Title)
	fmt.Println("price: ", resultBook1.Price)
	fmt.Println("rating: ", resultBook1.Rating)

	 var updateBook books.Book
	 err = db.Where("id = ?", 1).First(&updateBook).Error
	 if err != nil {
		fmt.Println("========================")
		fmt.Println("Error finding book record")
		fmt.Println("========================")
	 }

	 updateBook.Title = "Manusia Kuat (revised)"
	 err = db.Save(&updateBook).Error
	 if err != nil {
		fmt.Println("========================")
		fmt.Println("Error finding book record")
		fmt.Println("========================")
	 }
	 
	 var deleteBook books.Book
	 err = db.Where("id = ?", 3).Delete(&deleteBook).Error
	 if err != nil {
		fmt.Println("========================")
		fmt.Println("Error finding book record")
		fmt.Println("========================")
	 }


	router := gin.Default()

	// Versioning routes with group
	v1 := router.Group("/v1")

	// GET ROUTE
	v1.GET("/", handler.RootHandler)

	v1.GET("/book/:bookId", handler.BooksHandler)

	v1.GET("/book/:bookId/author/:authorId", handler.AuthorsHandler)

	v1.GET("/query", handler.QueryHandler)

	// POST ROUTE
	v1.POST("/book", handler.PostBookHandler)

	router.Run(":3000")
}
