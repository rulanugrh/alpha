package main

import (
	booksControll "github.com/rulanugrh/alpha/apigateway/controller/book"
	routers "github.com/rulanugrh/alpha/apigateway/router"
	booksService "github.com/rulanugrh/alpha/apigateway/services/book"
)

func main() {
	bookServ := booksService.NewBookServices()
	bookControll := booksControll.NewBookController(bookServ)

	routers.App(bookControll)

}
