package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/danzBraham/eniqilo-store/internal/entities/userentity"
	"github.com/danzBraham/eniqilo-store/internal/errors/usererror"
	"github.com/danzBraham/eniqilo-store/internal/helpers/httphelper"
	"github.com/danzBraham/eniqilo-store/internal/services"
)

type UserController interface {
	HandleRegisterStaff(w http.ResponseWriter, r *http.Request)
	HandleLoginStaff(w http.ResponseWriter, r *http.Request)
	HandleRegisterCustomer(w http.ResponseWriter, r *http.Request)
}

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (c *UserControllerImpl) HandleRegisterStaff(w http.ResponseWriter, r *http.Request) {
	payload := &userentity.RegisterStaffRequest{}
	err := httphelper.DecodeAndValidate(w, r, payload)
	if err != nil {
		return
	}

	staffResponse, err := c.UserService.RegisterStaff(r.Context(), payload)
	if errors.Is(err, usererror.ErrPhoneNumberAlreadyExists) {
		httphelper.ErrorResponse(w, http.StatusConflict, err)
		return
	}
	if err != nil {
		httphelper.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	cookie := &http.Cookie{
		Name:    "Authorization",
		Value:   staffResponse.AccessToken,
		Expires: time.Now().Add(2 * time.Hour),
	}
	http.SetCookie(w, cookie)

	httphelper.SuccessResponse(w, http.StatusCreated, "User successfully registered", staffResponse)
}

func (c *UserControllerImpl) HandleLoginStaff(w http.ResponseWriter, r *http.Request) {
	payload := &userentity.LoginStaffRequest{}
	err := httphelper.DecodeAndValidate(w, r, payload)
	if err != nil {
		return
	}

	staffResponse, err := c.UserService.LoginStaff(r.Context(), payload)
	if errors.Is(err, usererror.ErrUserNotFound) {
		httphelper.ErrorResponse(w, http.StatusNotFound, err)
		return
	}
	if errors.Is(err, usererror.ErrInvalidPassword) {
		httphelper.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if err != nil {
		httphelper.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httphelper.SuccessResponse(w, http.StatusOK, "User successfully login", staffResponse)
}

func (c *UserControllerImpl) HandleRegisterCustomer(w http.ResponseWriter, r *http.Request) {
	payload := &userentity.RegisterCustomerRequest{}
	err := httphelper.DecodeAndValidate(w, r, payload)
	if err != nil {
		return
	}

	customerResponse, err := c.UserService.RegisterCustomer(r.Context(), payload)
	if errors.Is(err, usererror.ErrPhoneNumberAlreadyExists) {
		httphelper.ErrorResponse(w, http.StatusConflict, err)
		return
	}
	if err != nil {
		httphelper.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httphelper.SuccessResponse(w, http.StatusCreated, "User registered successfully", customerResponse)
}
