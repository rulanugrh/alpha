package main

import (
	booksControll "github.com/rulanugrh/alpha/book/handlers/book"
	cateControll "github.com/rulanugrh/alpha/book/handlers/category"
	booksRepo "github.com/rulanugrh/alpha/book/repository/book"
	cateRepo "github.com/rulanugrh/alpha/book/repository/category"
	"github.com/rulanugrh/alpha/book/routes"
	booksServ "github.com/rulanugrh/alpha/book/services/book"
	cateServ "github.com/rulanugrh/alpha/book/services/category"
)

func main() {
	bookRepository := booksRepo.NewBookRepository()
	booksServices := booksServ.NewBookRepository(bookRepository)
	booksController := booksControll.NewBookController(booksServices)

	categoryRepository := cateRepo.NewCategoryRepository()
	categoryServices := cateServ.NewCategoryServices(categoryRepository)
	categoryController := cateControll.NewCategoryController(categoryServices)

	routes.Run(booksController, categoryController)
}
