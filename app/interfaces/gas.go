package interfaces

import (
	"github.com/dealense7/product-app/app/models"
)

type GasRepository interface {
	FindAll() ([]models.Gas, error)
}

type GasService interface {
	GetAll() ([]models.Gas, error)
}
