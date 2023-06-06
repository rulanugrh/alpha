package main

import (
	booksControll "github.com/rulanugrh/alpha/apigateway/controller/book"
	ordersControll "github.com/rulanugrh/alpha/apigateway/controller/order"
	routers "github.com/rulanugrh/alpha/apigateway/router"
	booksService "github.com/rulanugrh/alpha/apigateway/services/book"
	ordersService "github.com/rulanugrh/alpha/apigateway/services/order"
)

func main() {
	bookServ := booksService.NewBookServices()
	bookControll := booksControll.NewBookController(bookServ)

	orderServ := ordersService.NewOrderSercives()
	orderControll := ordersControll.NewOrderContoller(orderServ)
	routers.App(bookControll, orderControll)

}
