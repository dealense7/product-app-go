package repositories

import (
	"github.com/dealense7/product-app/app/interfaces"
	"github.com/dealense7/product-app/app/models"
	"github.com/jmoiron/sqlx"
)

type MySQLProductRepository struct {
	db *sqlx.DB
}

func NewMySQLProductRepository(db *sqlx.DB) interfaces.ProductRepository {
	return &MySQLProductRepository{db: db}
}

func (r *MySQLProductRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	query := `SELECT id, name FROM products`
	err := r.db.Select(&products, query)
	return products, err
}
