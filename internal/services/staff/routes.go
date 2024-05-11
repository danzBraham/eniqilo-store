package staff

import (
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /staff/register", h.handleRegister)
	router.HandleFunc("POST /staff/login", h.handleLogin)
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {}
