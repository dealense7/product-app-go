package interfaces

import (
	"github.com/dealense7/product-app/app/models"
)

type CurrencyRepository interface {
	FindAll() ([]models.Currency, error)
}

type CurrencyService interface {
	GetAll() ([]models.Currency, error)
}
