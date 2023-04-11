package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/rulanugrh/alpha/order/config"
	"github.com/rulanugrh/alpha/order/entities/domain"
	"github.com/rulanugrh/alpha/order/handlers/book"
	"github.com/rulanugrh/alpha/order/handlers/order"
	"github.com/rulanugrh/alpha/order/handlers/user"
)

func Run(book book.BookController, user user.UserController, order order.OrderController) {
	conf := config.GetConnect()
	conf.AutoMigrate(&domain.Book{}, &domain.User{}, &domain.Order{}, &domain.OrderItem{})

	serv := echo.New()
	apiBook := serv.Group("/books")
	{
		apiBook.GET("/", book.FindAll)
		apiBook.GET("/:id", book.FindId)
		apiBook.POST("/", book.Create)
	}

	apiUser := serv.Group("/users")
	{
		apiUser.POST("/", user.Create)
	}

	apiOrder := serv.Group("/orders")
	{
		apiOrder.POST("/", order.CreateOrder)
		apiOrder.GET("/:id", order.FindId)
		apiOrder.GET("/getAll/:id", order.FindAll)
		apiOrder.POST("/cart", order.AddCart)
		apiOrder.GET("/cart/:id", order.ListCart)
		apiOrder.DELETE("/cart/:id", order.DeleteCart)
		apiOrder.GET("/cart/:id/notpaid", order.ListNotPaid)
		apiOrder.PUT("/checkout/:id", order.Checkout)
	}

	// confApp := config.GetConfig()
	serv.Start(":8000")
}
