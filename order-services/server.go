package main

import (
	bookHand "github.com/rulanugrh/alpha/order/handlers/book"
	orderHand "github.com/rulanugrh/alpha/order/handlers/order"
	userHand "github.com/rulanugrh/alpha/order/handlers/user"
	bookRepo "github.com/rulanugrh/alpha/order/repository/book"
	orderRepo "github.com/rulanugrh/alpha/order/repository/order"
	userRepo "github.com/rulanugrh/alpha/order/repository/user"
	"github.com/rulanugrh/alpha/order/routes"
	bookServ "github.com/rulanugrh/alpha/order/services/book"
	orderServ "github.com/rulanugrh/alpha/order/services/order"
	userServ "github.com/rulanugrh/alpha/order/services/user"
)

func main() {
	bookRepository := bookRepo.NewBookRepository()
	bookServices := bookServ.NewBookServices(bookRepository)
	bookHandlers := bookHand.NewBookController(bookServices)

	userRepository := userRepo.NewUserRepository()
	userServices := userServ.NewUserServices(userRepository)
	userHandlers := userHand.NewUserController(userServices)

	orderRepository := orderRepo.NewOrderRepository()
	orderServices := orderServ.NewOrderServices(orderRepository)
	orderHandlers := orderHand.NewOrderController(orderServices)
	routes.Run(bookHandlers, userHandlers, orderHandlers)
}
