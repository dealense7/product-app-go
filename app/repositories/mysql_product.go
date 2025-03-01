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
	query := `SELECT pi.id,
					   pi.display_name_ka AS name,
					   pi.brand_name AS brand,
					   pi.unit,
					   pi.unit_type,
					   JSON_ARRAYAGG(JSON_OBJECT('provider', s.name, 'current_price', pp.current_price, 'created_at', DATE_FORMAT(pp.created_at, '%Y-%m-%d %H:%i'))) AS price_info,
					   SUBSTRING_INDEX(GROUP_CONCAT(f.url ORDER BY f.id ASC), ',', 1) AS image
				FROM
					(
						SELECT pp.item_id
						FROM product_prices pp
							JOIN product_items pi2 ON pi2.id = pp.item_id AND pi2.has_image = TRUE
						WHERE pp.status = TRUE
						GROUP BY pp.item_id
						HAVING COUNT(*) > 4
						ORDER BY RAND()
						LIMIT 7
					) AS top_items
						JOIN product_items pi ON pi.id = top_items.item_id
						JOIN product_prices pp ON pi.id = pp.item_id AND pp.status = TRUE
						JOIN stores s ON s.id = pp.store_id
						JOIN files f on pi.id = f.fileable_id
				GROUP BY pi.id, pi.display_name_ka, pi.brand_name, pi.unit, pi.unit_type;`
	err := r.db.Select(&products, query)
	return products, err
}
