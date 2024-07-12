package repositories

import (
	"context"
	"time"

	"github.com/danzBraham/eniqilo-store/internal/entities/productentity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *productentity.Product) (string, error)
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
