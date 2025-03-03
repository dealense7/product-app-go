package repositories

import (
	"github.com/dealense7/product-app/app/interfaces"
	"github.com/dealense7/product-app/app/models"
	"github.com/jmoiron/sqlx"
)

type MySQLCurrencyRepository struct {
	db *sqlx.DB
}

func NewMySQLCurrencyRepository(db *sqlx.DB) interfaces.CurrencyRepository {
	return &MySQLCurrencyRepository{db: db}
}

func (r *MySQLCurrencyRepository) FindAll() ([]models.Currency, error) {
	var items []models.Currency
	query := `select cr.id, cr.buy_rate, cr.sell_rate, cp.name, cp.logo_url, c.code
				from currency_rates as cr
				join currency_providers cp on cp.id = cr.provider_id
				join currencies c on c.id = cr.currency_id
				where cr.status is true
				;`
	err := r.db.Select(&items, query)
	return items, err
}
