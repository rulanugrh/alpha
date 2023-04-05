package category

import "github.com/rulanugrh/alpha/book/entities/domain"

type CategoryRepository interface {
	Create(category domain.Category) (domain.Category, error)
	Delete(id uint) error
}
