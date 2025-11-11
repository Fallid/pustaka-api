package main

import (
	"fmt"
	"pustaka-api/api"
	"pustaka-api/config"
	"pustaka-api/exception"
	"pustaka-api/repository"
	"pustaka-api/service"
	"pustaka-api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration from .env
	cfg := config.LoadConfig()

	// Setup database with config
	db := config.SetupDatabase(cfg)

	// Initialize repository
	bookRepository := repository.NewBookRepository(db)
	userRepository := repository.NewUserRepository(db)

	// Initialize service
	bookService := service.NewBookService(bookRepository)
	userService := service.NewUserService(userRepository)

	// Initialize handler
	bookHandler := api.NewBookHandler(bookService)
	userHandler := api.NewUserHandler(userService)

	// Setup router
	router := gin.Default()

	// Use custom logger middleware
	router.Use(utils.Logger())

	// Use error handler middleware
	router.Use(exception.ErrorHandler())

	// Versioning routes with group
	v1 := router.Group("/v1")
	{
		// User routes
		v1.POST("/register", userHandler.RegisterHandler)

		// Book routes
		v1.POST("/book", bookHandler.PostBookHandler)
		v1.GET("/book", bookHandler.GetBooksHandler)
		v1.GET("/book/:bookId", bookHandler.GetBookByIdHandler)
		v1.PUT("/book/:bookId", bookHandler.UpdateBookHandler)
		v1.DELETE("/book/:bookId", bookHandler.DeleteBookHandler)
	}

	// Run server with port from config
	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	router.Run(serverAddr)
}
