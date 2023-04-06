package category

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rulanugrh/alpha/book/entities/domain"
	"github.com/rulanugrh/alpha/book/entities/web"
	"github.com/rulanugrh/alpha/book/services/category"
)

type categorycontroller struct {
	cate category.CategoryServices
}

func NewCategoryController(ct category.CategoryServices) CategoryController {
	return &categorycontroller{
		cate: ct,
	}
}

func (cate *categorycontroller) CreateCategory(ctx *gin.Context) {
	category := domain.Category{}
	ctx.Bind(&category)

	response, err := cate.cate.CreateCategory(category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.CategoryFailure{
			Code:   500,
			Status: "CANT CREATE CATEGORY",
		})
		return
	}

	ctx.JSON(http.StatusOK, web.CategorySuccess{
		Code:   200,
		Status: "SUCCESS CREATE CATEGORY",
		Data:   response,
	})
}

func (cate *categorycontroller) DeleteCategory(ctx *gin.Context) {
	getID := ctx.Param("id")
	id, _ := strconv.Atoi(getID)

	err := cate.cate.DeleteCategory(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.CategoryFailure{
			Code:   500,
			Status: "CANT DELETE CATEGORY",
		})
		return
	}

	ctx.JSON(http.StatusOK, web.CategorySuccess{
		Code:   200,
		Status: "SUCCESS DELETE CATEGORY",
		Data:   nil,
	})
}
