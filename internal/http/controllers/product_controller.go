package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/danzBraham/eniqilo-store/internal/entities/productentity"
	"github.com/danzBraham/eniqilo-store/internal/errors/producterror"
	"github.com/danzBraham/eniqilo-store/internal/helpers/httphelper"
	"github.com/danzBraham/eniqilo-store/internal/services"
	"github.com/go-chi/chi/v5"
)

type ProductController interface {
	HandleCreateProduct(w http.ResponseWriter, r *http.Request)
	HandleGetProducts(w http.ResponseWriter, r *http.Request)
	HandleGetProductsForCustomer(w http.ResponseWriter, r *http.Request)
	HandleUpdateProductByID(w http.ResponseWriter, r *http.Request)
	HandleDeleteProductByID(w http.ResponseWriter, r *http.Request)
}

type ProductControllerImpl struct {
	ProductService services.ProductService
}

func NewProductController(productService services.ProductService) ProductController {
	return &ProductControllerImpl{ProductService: productService}
}

func (c *ProductControllerImpl) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	payload := &productentity.CreateProductRequest{}
	err := httphelper.DecodeAndValidate(w, r, payload)
	if err != nil {
		return
	}

	productResponse, err := c.ProductService.CreateProduct(r.Context(), payload)
	if err != nil {
		httphelper.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httphelper.SuccessResponse(w, http.StatusCreated, "success", productResponse)
}

func (c *ProductControllerImpl) HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	params := &productentity.ProductQueryParams{
		ID:          query.Get("id"),
		Limit:       5,
		Offset:      0,
		Name:        query.Get("name"),
		SKU:         query.Get("sku"),
		Category:    productentity.Category(query.Get("category")),
		Price:       query.Get("price"),
		InStock:     query.Get("inStock"),
		IsAvailable: query.Get("isAvailable"),
		CreatedAt:   query.Get("createdAt"),
	}

	if limit := query.Get("limit"); limit != "" {
		params.Limit, _ = strconv.Atoi(limit)
	}

	if offset := query.Get("offset"); offset != "" {
		params.Offset, _ = strconv.Atoi(offset)
	}

	productResponses, err := c.ProductService.GetProducts(r.Context(), params)
	if err != nil {
		httphelper.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httphelper.SuccessResponse(w, http.StatusOK, "success", productResponses)
}

func (c *ProductControllerImpl) HandleGetProductsForCustomer(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	params := &productentity.ProductQueryParams{
		Limit:    5,
		Offset:   0,
		Name:     query.Get("name"),
		SKU:      query.Get("sku"),
		Category: productentity.Category(query.Get("category")),
		Price:    query.Get("price"),
		InStock:  query.Get("inStock"),
	}

	if limit := query.Get("limit"); limit != "" {
		params.Limit, _ = strconv.Atoi(limit)
	}

	if offset := query.Get("offset"); offset != "" {
		params.Offset, _ = strconv.Atoi(offset)
	}

	productForCustomerResponses, err := c.ProductService.GetProductsForCustomer(r.Context(), params)
	if err != nil {
		httphelper.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httphelper.SuccessResponse(w, http.StatusOK, "successfully get products", productForCustomerResponses)
}

func (c *ProductControllerImpl) HandleUpdateProductByID(w http.ResponseWriter, r *http.Request) {
	payload := &productentity.UpdateProductRequest{}
	err := httphelper.DecodeAndValidate(w, r, payload)
	if err != nil {
		return
	}

	productID := chi.URLParam(r, "id")
	err = c.ProductService.UpdateProductByID(r.Context(), productID, payload)
	if errors.Is(err, producterror.ErrProductIDNotFound) {
		httphelper.ErrorResponse(w, http.StatusNotFound, err)
		return
	}
	if err != nil {
		httphelper.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httphelper.SuccessResponse(w, http.StatusOK, "successfully edit product", nil)
}

func (c *ProductControllerImpl) HandleDeleteProductByID(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "id")
	err := c.ProductService.DeleteProductByID(r.Context(), productID)
	if errors.Is(err, producterror.ErrProductIDNotFound) {
		httphelper.ErrorResponse(w, http.StatusNotFound, err)
		return
	}
	if err != nil {
		httphelper.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httphelper.SuccessResponse(w, http.StatusOK, "successfully delete product", nil)
}
