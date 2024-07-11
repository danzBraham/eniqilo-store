package http

import (
	"net/http"

	"github.com/danzBraham/eniqilo-store/internal/errors/commonerror"
	"github.com/danzBraham/eniqilo-store/internal/helpers/httphelper"
	"github.com/danzBraham/eniqilo-store/internal/http/controllers"
	"github.com/danzBraham/eniqilo-store/internal/repositories"
	"github.com/danzBraham/eniqilo-store/internal/services"
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

	// repositories
	userRepository := repositories.NewUserRepository(s.DB)

	// services
	userService := services.NewUserService(userRepository)

	// controllers
	userController := controllers.NewUserController(userService)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/staff", func(r chi.Router) {
			r.Post("/register", userController.HandleRegisterStaff)
			r.Post("/login", userController.HandleLoginStaff)
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
