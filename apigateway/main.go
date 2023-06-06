package main

import (
	booksControll "github.com/rulanugrh/alpha/apigateway/controller/book"
	ordersControll "github.com/rulanugrh/alpha/apigateway/controller/order"
	usersControll "github.com/rulanugrh/alpha/apigateway/controller/user"
	routers "github.com/rulanugrh/alpha/apigateway/router"
	booksService "github.com/rulanugrh/alpha/apigateway/services/book"
	ordersService "github.com/rulanugrh/alpha/apigateway/services/order"
	usersService "github.com/rulanugrh/alpha/apigateway/services/user"
)

func main() {
	bookServ := booksService.NewBookServices()
	bookControll := booksControll.NewBookController(bookServ)

	orderServ := ordersService.NewOrderSercives()
	orderControll := ordersControll.NewOrderContoller(orderServ)

	userServ := usersService.NewUserServices()
	userControll := usersControll.NewUserController(userServ)
	routers.App(bookControll, orderControll, userControll)

}
