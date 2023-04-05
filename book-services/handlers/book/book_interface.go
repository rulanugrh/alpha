package book

import "github.com/gin-gonic/gin"

type BookController interface {
	UploadBook(ctx *gin.Context)
	UpdateBook(ctx *gin.Context)
	DeleteBook(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindID(ctx *gin.Context)
	Seller(ctx *gin.Context)
}
