package services

import (
	"github.com/dealense7/product-app/app/interfaces"
	"github.com/dealense7/product-app/app/models"
)

type CurrencyService struct {
	repo interfaces.CurrencyRepository
}

func NewCurrencyService(repo interfaces.CurrencyRepository) interfaces.CurrencyService {
	return &CurrencyService{repo: repo}
}

func (s *CurrencyService) GetAll() ([]models.Currency, error) {
	return s.repo.FindAll()
}
