package category

import (
	"errors"

	"github.com/rulanugrh/alpha/book/config"
	"github.com/rulanugrh/alpha/book/entities/domain"
)

type categoryrepository struct{}

func NewCategoryRepository() CategoryRepository {
	return &categoryrepository{}
}

func (cate *categoryrepository) Create(category domain.Category) (domain.Category, error) {
	err := config.DB.Create(&category).Error
	if err != nil {
		return category, errors.New("cant create category")
	}

	return category, nil
}

func (cate *categoryrepository) Delete(id uint) error {
	var category domain.Category
	err := config.DB.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		panic(err)
	}

	return nil
}
