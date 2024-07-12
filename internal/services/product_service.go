package services

import (
	"context"

	"github.com/danzBraham/eniqilo-store/internal/entities/productentity"
	"github.com/danzBraham/eniqilo-store/internal/repositories"
	"github.com/oklog/ulid/v2"
)

type ProductService interface {
	CreateProduct(ctx context.Context, payload *productentity.CreateProductRequest) (*productentity.CreateProductResponse, error)
	GetProducts(ctx context.Context, params *productentity.ProductQueryParams) ([]*productentity.GetProductResponse, error)
}

type ProductServiceImpl struct {
	ProductRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) ProductService {
	return &ProductServiceImpl{ProductRepository: productRepository}
}

func (s *ProductServiceImpl) CreateProduct(ctx context.Context, payload *productentity.CreateProductRequest) (*productentity.CreateProductResponse, error) {
	product := &productentity.Product{
		ID:          ulid.Make().String(),
		Name:        payload.Name,
		SKU:         payload.SKU,
		Category:    payload.Category,
		ImageURL:    payload.ImageURL,
		Notes:       payload.Notes,
		Price:       payload.Price,
		Stock:       payload.Stock,
		Location:    payload.Location,
		IsAvailable: payload.IsAvailable,
	}

	createdAt, err := s.ProductRepository.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return &productentity.CreateProductResponse{
		ID:        product.ID,
		CreatedAt: createdAt,
	}, nil
}

func (s *ProductServiceImpl) GetProducts(ctx context.Context, params *productentity.ProductQueryParams) ([]*productentity.GetProductResponse, error) {
	return s.ProductRepository.GetProducts(ctx, params)
}
