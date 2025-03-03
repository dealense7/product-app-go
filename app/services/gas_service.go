package services

import (
	"github.com/dealense7/product-app/app/interfaces"
	"github.com/dealense7/product-app/app/models"
)

type GasService struct {
	repo interfaces.GasRepository
}

func NewGasService(repo interfaces.GasRepository) interfaces.GasService {
	return &GasService{repo: repo}
}

func (s *GasService) GetAll() ([]models.Gas, error) {
	return s.repo.FindAll()
}
