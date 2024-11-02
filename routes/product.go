package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/negeek/ecommerce-api-assessment/controllers"
	"github.com/negeek/ecommerce-api-assessment/middlewares"
)

func ProductRoutes(router *mux.Router, productController *controllers.ProductController) {
	authRoutes := router.PathPrefix("/products").Subrouter()
	authRoutes.Use(middlewares.AuthenticationMiddleware)
	authRoutes.HandleFunc("", productController.List).Methods(http.MethodGet)
	authRoutes.HandleFunc("/{id:[0-9]+}", productController.Find).Methods(http.MethodGet)

	adminRoutes := router.PathPrefix("/products").Subrouter()
	adminRoutes.Use(middlewares.AuthAdminMiddleware)

	adminRoutes.HandleFunc("", productController.Create).Methods(http.MethodPost)
	adminRoutes.HandleFunc("/{id:[0-9]+}", productController.Update).Methods(http.MethodPatch)
	adminRoutes.HandleFunc("/{id:[0-9]+}", productController.Delete).Methods(http.MethodDelete)
}
