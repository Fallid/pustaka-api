package main

import (

	"github.com/gin-gonic/gin"
	"pustaka-api/handler"
)

func main() {
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

