package main

import (
	bookHand "github.com/rulanugrh/alpha/order/handlers/book"
	bookRepo "github.com/rulanugrh/alpha/order/repository/book"
	"github.com/rulanugrh/alpha/order/routes"
	bookServ "github.com/rulanugrh/alpha/order/services/book"
)

func main() {
	bookRepository := bookRepo.NewBookRepository()
	bookServices := bookServ.NewBookServices(bookRepository)
	bookHandlers := bookHand.NewBookController(bookServices)

	routes.Run(bookHandlers)
}
