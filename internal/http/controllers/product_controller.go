package controllers

import (
	"net/http"

	"github.com/danzBraham/eniqilo-store/internal/entities/productentity"
	"github.com/danzBraham/eniqilo-store/internal/helpers/httphelper"
	"github.com/danzBraham/eniqilo-store/internal/services"
)

type ProductController interface {
	HandleCreateProduct(w http.ResponseWriter, r *http.Request)
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
