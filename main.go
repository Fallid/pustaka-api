package main

import (
	"fmt"
	"pustaka-api/api"
	"pustaka-api/config"
	"pustaka-api/exception"
	"pustaka-api/middleware"
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
	authenticationRepository := repository.NewAuthenticationRepository(db)

	// Initialize service
	bookService := service.NewBookService(bookRepository)
	userService := service.NewUserService(userRepository, authenticationRepository)

	// Initialize handler
	bookHandler := api.NewBookHandler(bookService)
	userHandler := api.NewUserHandler(userService, cfg)

	// Setup router
	router := gin.Default()

	// Use custom logger middleware
	router.Use(utils.Logger())

	// Use error handler middleware
	router.Use(exception.ErrorHandler())

	// Versioning routes with group
	v1 := router.Group("/v1")
	{
		// Public user routes
		v1.POST("/register", userHandler.RegisterHandler)
		v1.POST("/login", userHandler.LoginHandler)
		v1.POST("/refresh-token", userHandler.RefreshTokenHandler)
		v1.POST("/logout", userHandler.LogoutHandler)
		v1.GET("/users/:userId", userHandler.GetUserByIdHandler)

		// Public book routes (read only)
		
		// Protected book routes (require authentication)
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware(cfg))
		{
			protected.GET("/book", bookHandler.GetBooksHandler)
			protected.GET("/book/:bookId", bookHandler.GetBookByIdHandler)
			protected.POST("/book", bookHandler.PostBookHandler)
			protected.PUT("/book/:bookId", bookHandler.UpdateBookHandler)
			protected.DELETE("/book/:bookId", bookHandler.DeleteBookHandler)
		}
	}

	// Run server with port from config
	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	router.Run(serverAddr)
}
