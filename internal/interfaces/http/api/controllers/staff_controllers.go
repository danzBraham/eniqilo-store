package controllers

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/danzBraham/eniqilo-store/internal/applications/interfaces"
	"github.com/danzBraham/eniqilo-store/internal/domains/entities"
	"github.com/danzBraham/eniqilo-store/internal/helpers"
)

type StaffController struct {
	StaffService interfaces.StaffService
	Router       *http.ServeMux
}

func NewStaffController(staffService interfaces.StaffService, router *http.ServeMux) *StaffController {
	controller := &StaffController{
		StaffService: staffService,
		Router:       router,
	}

	controller.Router.HandleFunc("POST /staff/register", controller.handleRegister)
	controller.Router.HandleFunc("POST /staff/login", controller.handleLogin)

	return controller
}

func (s *StaffController) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Get JSON payload
	payload := &entities.RegisterStaff{}
	if err := helpers.ParseJSON(r, payload); err != nil {
		helpers.ResponseJSON(w, http.StatusBadRequest, &ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	// Validate request payload
	if err := helpers.ValidatePayload(payload); err != nil {
		helpers.ResponseJSON(w, http.StatusBadRequest, &ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	// Validate phone number
	isValid, err := helpers.ValidatePhoneNumber(payload.PhoneNumber)
	if err != nil {
		helpers.ResponseJSON(w, http.StatusInternalServerError, &ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if !isValid {
		helpers.ResponseJSON(w, http.StatusBadRequest, &ErrorResponse{
			Message: "invalid phone number",
		})
		return
	}

	staff, err := s.StaffService.RegisterStaff(payload)
	if err != nil {
		helpers.ResponseJSON(w, http.StatusInternalServerError, &ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	jwtExp, _ := strconv.Atoi(os.Getenv("JWT_EXP"))

	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    staff.AccessToken,
		Expires:  time.Now().Add(time.Duration(jwtExp) * time.Hour),
		HttpOnly: true,
		Secure:   false,
	}
	http.SetCookie(w, cookie)

	helpers.ResponseJSON(w, http.StatusCreated, &SuccessResponse{
		Message: "User successfully registered",
		Data:    staff,
	})
}

func (s *StaffController) handleLogin(w http.ResponseWriter, r *http.Request) {
	// Get JSON payload
	payload := &entities.LoginStaff{}
	if err := helpers.ParseJSON(r, payload); err != nil {
		helpers.ResponseJSON(w, http.StatusBadRequest, &ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	// Validate request payload
	if err := helpers.ValidatePayload(payload); err != nil {
		helpers.ResponseJSON(w, http.StatusBadRequest, &ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	// Validate phone number
	isValid, err := helpers.ValidatePhoneNumber(payload.PhoneNumber)
	if err != nil {
		helpers.ResponseJSON(w, http.StatusInternalServerError, &ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if !isValid {
		helpers.ResponseJSON(w, http.StatusBadRequest, &ErrorResponse{
			Message: "invalid phone number",
		})
		return
	}

	staff, err := s.StaffService.LoginStaff(payload)
	if err != nil {
		helpers.ResponseJSON(w, http.StatusInternalServerError, &ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	jwtExp, _ := strconv.Atoi(os.Getenv("JWT_EXP"))

	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    staff.AccessToken,
		Expires:  time.Now().Add(time.Duration(jwtExp) * time.Hour),
		HttpOnly: true,
		Secure:   false,
	}
	http.SetCookie(w, cookie)

	helpers.ResponseJSON(w, http.StatusCreated, &SuccessResponse{
		Message: "User successfully login",
		Data:    staff,
	})
}
