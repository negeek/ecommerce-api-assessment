package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/negeek/ecommerce-api-assessment/controllers"
)

func UserRoutes(router *mux.Router, userController *controllers.UserController) {
	userRoutes := router.PathPrefix("/users").Subrouter()
	userRoutes.HandleFunc("/register", userController.Register).Methods(http.MethodPost)
	userRoutes.HandleFunc("/login", userController.Login).Methods(http.MethodPost)
}
