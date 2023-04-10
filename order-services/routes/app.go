package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/rulanugrh/alpha/order/config"
	"github.com/rulanugrh/alpha/order/entities/domain"
	"github.com/rulanugrh/alpha/order/handlers/book"
	"github.com/rulanugrh/alpha/order/handlers/user"
)

func Run(book book.BookController, user user.UserController) {
	conf := config.GetConnect()
	conf.AutoMigrate(&domain.Book{}, &domain.User{}, &domain.Order{}, &domain.OrderItem{})

	serv := echo.New()
	apiBook := serv.Group("/books/")
	{
		apiBook.GET("/", book.FindAll)
		apiBook.GET("/:id", book.FindId)
		apiBook.POST("/", book.Create)
	}

	apiUser := serv.Group("/users/")
	{
		apiUser.POST("/", user.Create)
	}

	confApp := config.GetConfig()
	serv.Start(confApp.Host + ":" + confApp.Port)
}
