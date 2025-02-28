package interfaces

import (
	"github.com/dealense7/product-app/app/models"
)

type ProductRepository interface {
	FindAll() ([]models.Product, error)
}

type ProductService interface {
	GetProducts() ([]models.Product, error)
}
