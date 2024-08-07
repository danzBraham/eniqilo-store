package http

import (
	"net/http"

	"github.com/danzBraham/eniqilo-store/internal/errors/commonerror"
	"github.com/danzBraham/eniqilo-store/internal/helpers/httphelper"
	_ "github.com/danzBraham/eniqilo-store/internal/helpers/validator"
	"github.com/danzBraham/eniqilo-store/internal/http/controllers"
	"github.com/danzBraham/eniqilo-store/internal/http/middlewares"
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
	productRepository := repositories.NewProductRepository(s.DB)
	checkoutRepository := repositories.NewCheckoutRepository(s.DB)

	// services
	userService := services.NewUserService(userRepository)
	productService := services.NewProductService(productRepository)
	checkoutService := services.NewCheckoutService(checkoutRepository, userRepository)

	// controllers
	userController := controllers.NewUserController(userService)
	productController := controllers.NewProductController(productService)
	checkoutController := controllers.NewCheckoutController(checkoutService)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/staff", func(r chi.Router) {
			r.Post("/register", userController.HandleRegisterStaff)
			r.Post("/login", userController.HandleLoginStaff)
		})

		r.Get("/product/customer", productController.HandleGetProductsForCustomer)

		r.Group(func(r chi.Router) {
			r.Use(middlewares.Auth)

			r.Route("/product", func(r chi.Router) {
				r.Post("/", productController.HandleCreateProduct)
				r.Get("/", productController.HandleGetProducts)
				r.Put("/{id}", productController.HandleUpdateProductByID)
				r.Delete("/{id}", productController.HandleDeleteProductByID)

				r.Route("/checkout", func(r chi.Router) {
					r.Post("/", checkoutController.HandleCheckoutProduct)
					r.Get("/history", checkoutController.HandleGetCheckoutHistroies)
				})
			})

			r.Route("/customer", func(r chi.Router) {
				r.Post("/register", userController.HandleRegisterCustomer)
				r.Get("/", userController.HandleGetCustomers)
			})
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
