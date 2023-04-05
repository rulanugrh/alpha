package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rulanugrh/alpha/book/config"
	"github.com/rulanugrh/alpha/book/entities/domain"
	books "github.com/rulanugrh/alpha/book/handlers/book"
	"github.com/rulanugrh/alpha/book/handlers/category"
)

func Run(book books.BookController, category category.CategoryController) {
	db := config.GetConnection()
	db.AutoMigrate(&domain.Category{}, &domain.Seller{}, &domain.Book{})

	server := gin.Default()

	bookGroup := server.Group("/api/book")
	{
		bookGroup.GET("/", book.FindAll)
		bookGroup.POST("/", book.UploadBook)
		bookGroup.PUT("/:id", book.UpdateBook)
		bookGroup.GET("/:id", book.FindID)
		bookGroup.DELETE("/:id", book.DeleteBook)
	}

	categoryGroup := server.Group("/api/categories")
	{
		categoryGroup.POST("/", category.CreateCategory)
		categoryGroup.DELETE("/:id", category.DeleteCategory)
	}

	sellerGroup := server.Group("/api/seller")
	{
		sellerGroup.POST("/", book.Seller)
	}

	conf := config.GetConfig()

	server.Run(conf.RunnApp.Host + ":" + conf.RunnApp.Port)
}
