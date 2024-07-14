package controllers

import (
	"errors"
	"net/http"

	"github.com/danzBraham/eniqilo-store/internal/entities/checkoutentity"
	"github.com/danzBraham/eniqilo-store/internal/errors/checkouterror"
	"github.com/danzBraham/eniqilo-store/internal/helpers/httphelper"
	"github.com/danzBraham/eniqilo-store/internal/services"
)

type CheckoutController interface {
	HandleCheckoutProduct(w http.ResponseWriter, r *http.Request)
}

type CheckoutControllerImpl struct {
	CheckoutService services.CheckoutService
}

func NewCheckoutController(checkoutService services.CheckoutService) CheckoutController {
	return &CheckoutControllerImpl{CheckoutService: checkoutService}
}

func (c *CheckoutControllerImpl) HandleCheckoutProduct(w http.ResponseWriter, r *http.Request) {
	payload := &checkoutentity.CheckoutProductRequest{}
	err := httphelper.DecodeAndValidate(w, r, payload)
	if err != nil {
		return
	}

	err = c.CheckoutService.CheckoutProduct(r.Context(), payload)
	if errors.Is(err, checkouterror.ErrCustomerIDNotFound) {
		httphelper.ErrorResponse(w, http.StatusNotFound, err)
		return
	}
	if errors.Is(err, checkouterror.ErrProductIDNotFound) {
		httphelper.ErrorResponse(w, http.StatusNotFound, err)
		return
	}
	if errors.Is(err, checkouterror.ErrOneOfProductNotAvailable) {
		httphelper.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if errors.Is(err, checkouterror.ErrOneOfProductStockNotEnough) {
		httphelper.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if errors.Is(err, checkouterror.ErrPaidNotEnough) {
		httphelper.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if errors.Is(err, checkouterror.ErrChangeNotRight) {
		httphelper.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if err != nil {
		httphelper.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httphelper.SuccessResponse(w, http.StatusOK, "successfully checkout product", nil)
}
