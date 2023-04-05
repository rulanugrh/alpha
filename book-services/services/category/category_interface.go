package category

import (
	"github.com/rulanugrh/alpha/book/entities/domain"
	"github.com/rulanugrh/alpha/book/entities/web"
)

type CategoryServices interface {
	CreateCategory(category domain.Category) (web.CategoryResponse, error)
	DeleteCategory(id uint) error
}
