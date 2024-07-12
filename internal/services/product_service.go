package services

import (
	"context"

	"github.com/danzBraham/eniqilo-store/internal/entities/productentity"
	"github.com/danzBraham/eniqilo-store/internal/errors/producterror"
	"github.com/danzBraham/eniqilo-store/internal/repositories"
	"github.com/oklog/ulid/v2"
)

type ProductService interface {
	CreateProduct(ctx context.Context, payload *productentity.CreateProductRequest) (*productentity.CreateProductResponse, error)
	GetProducts(ctx context.Context, params *productentity.ProductQueryParams) ([]*productentity.GetProductResponse, error)
	GetProductsForCustomer(ctx context.Context, params *productentity.ProductQueryParams) ([]*productentity.GetProductResponse, error)
	UpdateProductByID(ctx context.Context, productID string, payload *productentity.UpdateProductRequest) error
	DeleteProductByID(ctx context.Context, productID string) error
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

func (s *ProductServiceImpl) GetProductsForCustomer(ctx context.Context, params *productentity.ProductQueryParams) ([]*productentity.GetProductResponse, error) {
	return s.ProductRepository.GetProductsForCustomer(ctx, params)
}

func (s *ProductServiceImpl) UpdateProductByID(ctx context.Context, productID string, payload *productentity.UpdateProductRequest) error {
	IsProductIDExists, err := s.ProductRepository.IsProductIDExists(ctx, productID)
	if err != nil {
		return err
	}
	if !IsProductIDExists {
		return producterror.ErrProductIDNotFound
	}

	product := &productentity.Product{
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

	err = s.ProductRepository.UpdateProductByID(ctx, productID, product)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductServiceImpl) DeleteProductByID(ctx context.Context, productID string) error {
	IsProductIDExists, err := s.ProductRepository.IsProductIDExists(ctx, productID)
	if err != nil {
		return err
	}
	if !IsProductIDExists {
		return producterror.ErrProductIDNotFound
	}

	err = s.ProductRepository.DeleteProductByID(ctx, productID)
	if err != nil {
		return err
	}

	return nil
}
