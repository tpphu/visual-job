package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/tpphu/visual-job/model"
)

// ProductRepo interface
type ProductRepo interface {
	Create(model.Product) (*model.Product, error)
	List(int, int) ([]model.Product, error)
}

// ProductRepoImpl struct
type ProductRepoImpl struct {
	DB *gorm.DB
}

// Create returns a product
func (productRepo ProductRepoImpl) Create(product model.Product) (*model.Product, error) {
	err := productRepo.DB.Create(&product).Error
	return &product, err
}

func (productRepo ProductRepoImpl) List(offsetID int, limit int) ([]model.Product, error) {
	if limit == 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	products := []model.Product{}
	err := productRepo.DB.
		Where("id > ?", offsetID).
		Where("id < 1000").
		Limit(limit).
		Find(&products).
		Error
	return products, err
}
