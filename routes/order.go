package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/negeek/ecommerce-api-assessment/controllers"
	"github.com/negeek/ecommerce-api-assessment/middlewares"
)

func OrderRoutes(router *mux.Router, ordersController *controllers.OrdersController) {
	authRoutes := router.PathPrefix("/orders").Subrouter()
	authRoutes.Use(middlewares.AuthenticationMiddleware)
	authRoutes.HandleFunc("", ordersController.PlaceOrder).Methods(http.MethodPost)
	authRoutes.HandleFunc("", ordersController.ListOrders).Methods(http.MethodGet)
	authRoutes.HandleFunc("/{id:[0-9]+}/cancel", ordersController.CancelOrder).Methods(http.MethodPatch)
	adminRoutes := router.PathPrefix("/orders").Subrouter()
	adminRoutes.Use(middlewares.AuthAdminMiddleware)
	adminRoutes.HandleFunc("/{id:[0-9]+}/update-status", ordersController.UpdateOrderStatus).Methods(http.MethodPatch)
}
