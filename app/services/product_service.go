package services

import (
	"github.com/dealense7/product-app/app/interfaces"
	"github.com/dealense7/product-app/app/models"
)

type ProductService struct {
	repo interfaces.ProductRepository
}

func NewProductService(repo interfaces.ProductRepository) interfaces.ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProducts(filters map[string]interface{}) ([]models.Product, error) {
	return s.repo.FindAll(filters)
}

func (s *ProductService) GroupByCategory() ([]models.Category, error) {
	return s.repo.GroupByCategory()
}
