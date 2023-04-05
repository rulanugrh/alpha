package category

import (
	"errors"

	"github.com/rulanugrh/alpha/book/entities/domain"
	"github.com/rulanugrh/alpha/book/entities/web"
	categories "github.com/rulanugrh/alpha/book/repository/category"
)

type categoryservices struct {
	cate categories.CategoryRepository
}

func NewCategoryServices(cates categories.CategoryRepository) CategoryServices {
	return &categoryservices{
		cate: cates,
	}
}

func (cate *categoryservices) CreateCategory(category domain.Category) (web.CategoryResponse, error) {
	response, err := cate.cate.Create(category)
	if err != nil {
		return web.CategoryResponse{}, errors.New("cant create category")
	}

	return web.CategoryResponse{
		Name:        response.Name,
		Description: response.Description,
	}, nil
}

func (cate *categoryservices) DeleteCategory(id uint) error {
	err := cate.cate.Delete(id)
	if err != nil {
		return errors.New("cant delete category")
	}

	return nil
}
