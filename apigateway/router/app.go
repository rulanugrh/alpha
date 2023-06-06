package routers

import (
	"github.com/labstack/echo/v4"
	bookControll "github.com/rulanugrh/alpha/apigateway/controller/book"
	orderControll "github.com/rulanugrh/alpha/apigateway/controller/order"
	userControll "github.com/rulanugrh/alpha/apigateway/controller/user"
	"github.com/rulanugrh/alpha/apigateway/helper"
)

func App(books bookControll.BookController, orders orderControll.OrderInterfaces, users userControll.UserController) {

	serv := echo.New()
	apiBook := serv.Group("/book")
	{
		apiBook.GET("/", books.GetBook)
		apiBook.GET("/:id", books.GetBookId)
		apiBook.POST("/", books.PostBook)
		apiBook.PUT("/:id", books.UpdateBook)
		apiBook.DELETE("/:id", books.DeleteBookById)
	}

	apiCate := serv.Group("/category")
	{
		apiCate.GET("/", books.GetCategory)
		apiCate.POST("/", books.GetCategory)
		apiCate.GET("/:id", books.GetCategoryId)
	}

	apiOrder := serv.Group("/order")
	{
		apiOrder.POST("/", orders.Create)
		apiOrder.GET("/:id", orders.FindId)
		apiOrder.GET("/getAll/:id", orders.FindAll)
		apiOrder.POST("/cart/", orders.AddCart)
		apiOrder.PUT("/checkout/:id", orders.Checkout)
		apiOrder.GET("/cart/:userId", orders.ListCart)
		apiOrder.DELETE("/cart/:id", orders.DeleteCart)
		apiOrder.GET("/cart/:id/notpaid", orders.ListNotPaid)
	}

	apiUser := serv.Group("/users")
	{
		apiUser.POST("/register", users.Register)
		apiUser.POST("/login", users.Login)
		apiUser.GET("/detail/:id", users.Detail)
	}

	err := serv.Start(":8080")
	if err != nil {
		helper.PanicIfError(err)
	}
}
