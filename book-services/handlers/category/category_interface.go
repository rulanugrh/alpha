package category

import "github.com/gin-gonic/gin"

type CategoryController interface {
	CreateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
}
