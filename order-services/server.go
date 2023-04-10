package main

import (
	bookHand "github.com/rulanugrh/alpha/order/handlers/book"
	userHand "github.com/rulanugrh/alpha/order/handlers/user"
	bookRepo "github.com/rulanugrh/alpha/order/repository/book"
	userRepo "github.com/rulanugrh/alpha/order/repository/user"
	"github.com/rulanugrh/alpha/order/routes"
	bookServ "github.com/rulanugrh/alpha/order/services/book"
	userServ "github.com/rulanugrh/alpha/order/services/user"
)

func main() {
	bookRepository := bookRepo.NewBookRepository()
	bookServices := bookServ.NewBookServices(bookRepository)
	bookHandlers := bookHand.NewBookController(bookServices)

	userRepository := userRepo.NewUserRepository()
	userServices := userServ.NewUserServices(userRepository)
	userHandlers := userHand.NewUserController(userServices)

	routes.Run(bookHandlers, userHandlers)
}
