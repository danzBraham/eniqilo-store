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

	staffResponse, err := c.UserService.CreateStaff(r.Context(), payload)
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
