package interfaces

import (
	"github.com/dealense7/product-app/app/models"
)

type ProductRepository interface {
	FindAll(filters map[string]interface{}) ([]models.Product, error)
	GetFilteredProducts(filters map[string]interface{}) ([]models.Product, error)
	GroupByCategory() ([]models.Category, error)
}

type ProductService interface {
	GetProducts(map[string]interface{}) ([]models.Product, error)
	GroupByCategory() ([]models.Category, error)
}
