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
	// db.AutoMigrate(books.Book{})

	// repository
	bookRepository := books.NewRepository(db)
	// service
	bookService := books.NewService(bookRepository)
	// handler
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	// Versioning routes with group
	v1 := router.Group("/v1")

	// GET ROUTE
	// v1.GET("/", bookHandler.RootHandler)

	// v1.GET("/book/:bookId", handler.BooksHandler)

	// v1.GET("/book/:bookId/author/:authorId", handler.AuthorsHandler)

	// v1.GET("/query", handler.QueryHandler)

	// POST ROUTE
	v1.POST("/book", bookHandler.PostBookHandler)
	v1.GET("/book", bookHandler.GetBooksHandler)
	v1.GET("/book/:bookId", bookHandler.GetBookByIdHandler)
	v1.PUT("/book/:bookId", bookHandler.UpdateBookHandler)
	v1.DELETE("/book/:bookId", bookHandler.DeleteBookHandler)
	router.Run(":3000")
}
