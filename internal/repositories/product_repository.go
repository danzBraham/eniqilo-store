package repositories

import (
	"context"
	"strconv"
	"time"

	"github.com/danzBraham/eniqilo-store/internal/entities/productentity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *productentity.Product) (string, error)
	GetProducts(ctx context.Context, params *productentity.ProductQueryParams) ([]*productentity.GetProductResponse, error)
}

type ProductRepositoryImpl struct {
	DB *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) ProductRepository {
	return &ProductRepositoryImpl{DB: db}
}

func (r *ProductRepositoryImpl) CreateProduct(ctx context.Context, product *productentity.Product) (string, error) {
	query := `
		INSERT INTO
			products (id, name, sku, category, image_url, notes, price, stock, location, is_available)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING
			created_at
	`
	var createdAt time.Time
	err := r.DB.QueryRow(ctx, query,
		&product.ID,
		&product.Name,
		&product.SKU,
		&product.Category,
		&product.ImageURL,
		&product.Notes,
		&product.Price,
		&product.Stock,
		&product.Location,
		&product.IsAvailable,
	).Scan(&createdAt)
	if err != nil {
		return "", err
	}
	return createdAt.Format(time.RFC3339), nil
}

func (r *ProductRepositoryImpl) GetProducts(ctx context.Context, params *productentity.ProductQueryParams) ([]*productentity.GetProductResponse, error) {
	query := `
		SELECT
			id,
			name,
			sku,
			category,
			image_url,
			notes,
			price,
			stock,
			location,
			is_available,
			created_at
		FROM
			products
		WHERE
			is_deleted = false
	`
	args := []interface{}{}
	argID := 1

	if params.ID != "" {
		query += ` AND id = $` + strconv.Itoa(argID)
		args = append(args, params.ID)
		argID++
	}

	if params.Name != "" {
		query += ` AND name ILIKE $` + strconv.Itoa(argID)
		args = append(args, `%`+params.Name+`%`)
		argID++
	}

	if params.SKU != "" {
		query += ` AND sku = $` + strconv.Itoa(argID)
		args = append(args, params.SKU)
		argID++
	}

	validCategories := map[productentity.Category]struct{}{
		productentity.Clothing:    {},
		productentity.Accessories: {},
		productentity.Footwear:    {},
		productentity.Beverages:   {},
	}

	if _, ok := validCategories[params.Category]; ok {
		query += ` AND category = $` + strconv.Itoa(argID)
		args = append(args, params.Category)
		argID++
	}

	switch params.InStock {
	case "true":
		query += ` AND stock > 0`
	case "false":
		query += ` AND stock = 0`
	}

	switch params.IsAvailable {
	case "true":
		query += ` AND is_available = true`
	case "false":
		query += ` AND is_available = false`
	}

	switch params.Price {
	case "asc":
		query += ` ORDER BY price ASC`
	case "desc":
		query += ` ORDER BY price DESC`
	}

	switch params.CreatedAt {
	case "asc":
		query += ` ORDER BY created_at ASC`
	case "desc":
		query += ` ORDER BY created_at DESC`
	}

	query += ` LIMIT $` + strconv.Itoa(argID) + ` OFFSET $` + strconv.Itoa(argID+1)
	args = append(args, params.Limit, params.Offset)

	rows, err := r.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]*productentity.GetProductResponse, 0, params.Limit)
	for rows.Next() {
		var product productentity.GetProductResponse
		var createdAt time.Time
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.SKU,
			&product.Category,
			&product.ImageURL,
			&product.Notes,
			&product.Price,
			&product.Stock,
			&product.Location,
			&product.IsAvailable,
			&createdAt,
		)
		if err != nil {
			return nil, err
		}
		product.CreatedAt = createdAt.Format(time.RFC3339)
		products = append(products, &product)
	}

	return products, nil
}
