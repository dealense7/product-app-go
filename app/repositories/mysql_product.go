package repositories

import (
	"errors"
	"strings"

	"github.com/dealense7/product-app/app/interfaces"
	"github.com/dealense7/product-app/app/models"
	"github.com/jmoiron/sqlx"
)

// MySQLProductRepository implements the ProductRepository interface.
type MySQLProductRepository struct {
	db *sqlx.DB
}

// NewMySQLProductRepository creates a new instance of MySQLProductRepository.
func NewMySQLProductRepository(db *sqlx.DB) interfaces.ProductRepository {
	return &MySQLProductRepository{db: db}
}

// buildProductQuery builds the common product query.
// threshold: the HAVING COUNT(*) threshold (e.g. 4 for FindAll, 2 for GetFilteredProducts)
// extraCondition: additional filtering condition for the outer query (e.g. "pi.category_id = ?")
func buildProductQuery(threshold int, extraCondition string) string {
	baseQuery := `
		SELECT 
			pi.id,
			pi.display_name_ka AS name,
			pi.brand_name AS brand,
			pi.unit,
			pi.unit_type,
			JSON_ARRAYAGG(
				JSON_OBJECT(
					'provider', s.name, 
					'current_price', pp.current_price, 
					'created_at', DATE_FORMAT(pp.created_at, '%Y-%m-%d %H:%i')
				)
			) AS price_info,
			SUBSTRING_INDEX(GROUP_CONCAT(f.url ORDER BY f.id ASC), ',', 1) AS image
		FROM (
			SELECT pp.item_id
			FROM product_prices pp
			JOIN product_items pi2 ON pi2.id = pp.item_id AND pi2.has_image = TRUE
			WHERE pp.status = TRUE`
	// For the subquery, if we need to filter by category (as in GetFilteredProducts) we add it here.
	if extraCondition != "" && strings.Contains(extraCondition, "pi2.category_id") {
		baseQuery += " AND " + extraCondition
	}
	baseQuery += `
			GROUP BY pp.item_id
			HAVING COUNT(*) > ?
			ORDER BY RAND()
			LIMIT 7
		) AS top_items
		JOIN product_items pi ON pi.id = top_items.item_id
		JOIN product_prices pp ON pi.id = pp.item_id AND pp.status = TRUE
		JOIN stores s ON s.id = pp.store_id
		JOIN files f ON pi.id = f.fileable_id`
	// If extraCondition does not involve the subquery filtering, apply it in the outer query.
	if extraCondition != "" && !strings.Contains(extraCondition, "pi2.category_id") {
		baseQuery += " WHERE " + extraCondition
	}
	baseQuery += `
		GROUP BY pi.id, pi.display_name_ka, pi.brand_name, pi.unit, pi.unit_type`
	return baseQuery
}

// FindAll retrieves products with optional filtering.
func (r *MySQLProductRepository) FindAll(filters map[string]interface{}) ([]models.Product, error) {
	var products []models.Product
	var condition string
	var args []interface{}

	// Use type assertion for the filter value.
	if category, ok := filters["categoryId"].(string); ok && category != "" {
		condition = "pi.category_id = ?"
		args = append(args, category)
	}

	// Build the query; threshold is 4 in this case.
	query := buildProductQuery(4, condition)

	// Prepend the threshold argument for the subquery HAVING clause.
	args = append([]interface{}{4}, args...)

	err := r.db.Select(&products, query, args...)
	return products, err
}

// GetFilteredProducts retrieves products filtered by a category.
func (r *MySQLProductRepository) GetFilteredProducts(filters map[string]interface{}) ([]models.Product, error) {
	var products []models.Product

	category, ok := filters["categoryId"]
	if !ok {
		return products, errors.New("categoryId filter is required")
	}

	// For GetFilteredProducts, the category filter is applied inside the subquery.
	condition := "pi2.category_id = ?"
	query := buildProductQuery(2, condition)
	err := r.db.Select(&products, query, category, 2)
	return products, err
}

// GroupByCategory retrieves categories that meet criteria and attaches filtered products.
func (r *MySQLProductRepository) GroupByCategory() ([]models.Category, error) {
	var categories []models.Category
	query := `
		SELECT 
			c.id, 
			c.name
		FROM categories c
		JOIN product_items pi ON c.id = pi.category_id AND pi.has_image = TRUE
		GROUP BY c.id, c.name
		HAVING COUNT(pi.id) > 7
		ORDER BY RAND()
		LIMIT 3`
	err := r.db.Select(&categories, query)
	if err != nil {
		return nil, err
	}

	// Note: This loop performs additional queries (N+1). For better performance with many categories,
	// consider joining products to categories in one query.
	for i := range categories {
		filters := map[string]interface{}{
			"categoryId": categories[i].ID,
		}
		products, err := r.GetFilteredProducts(filters)
		if err != nil {
			return nil, err
		}
		categories[i].Products = products
	}
	return categories, nil
}
