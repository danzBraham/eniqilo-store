package http

import (
	"net/http"

	"github.com/danzBraham/eniqilo-store/internal/errors/commonerror"
	"github.com/danzBraham/eniqilo-store/internal/helpers/httphelper"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		httphelper.EncodeJSON(w, http.StatusOK, httphelper.ResponseBody{
			Message: "Welcome to Eniqilo Store API",
		})
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		httphelper.ErrorResponse(w, http.StatusNotFound, commonerror.ErrRouteDoesNotExist)
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		httphelper.ErrorResponse(w, http.StatusMethodNotAllowed, commonerror.ErrMethodNotAllowed)
	})

	return r
}
