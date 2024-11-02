package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/negeek/ecommerce-api-assessment/controllers"
	"github.com/negeek/ecommerce-api-assessment/db"
	"github.com/negeek/ecommerce-api-assessment/middlewares"
	"github.com/negeek/ecommerce-api-assessment/routes"
	"github.com/negeek/ecommerce-api-assessment/services"
)

func loadEnv() {
	log.Println(("load env"))
	environment := os.Getenv("ENVIRONMENT")
	if environment == "dev" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("error loading .env file: ", err)
		}
	}
}

func connectDB() {
	dbUrl := os.Getenv("POSTGRESQL_URL")
	log.Println("connecting to DB...")
	if err := db.Connect(dbUrl); err != nil {
		log.Fatal("failed to connect to DB: ", err)
	}
	log.Println("connected to DB")
}

func setupRouter() *mux.Router {
	log.Println("setup router")
	router := mux.NewRouter()
	router.Use(middlewares.CORS)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods(http.MethodGet)

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	userService := &services.UserService{}
	productService := &services.ProductService{}
	orderService := &services.OrderService{}

	userController := controllers.NewUserController(userService)
	productController := controllers.NewProductController(productService)
	ordersController := controllers.NewOrdersController(orderService)

	routes.OrderRoutes(apiRouter, ordersController)
	routes.ProductRoutes(apiRouter, productController)
	routes.UserRoutes(apiRouter, userController)

	return router
}

func startServer() *http.Server {
	router := setupRouter()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Println("start server...")
		log.Println("access it at http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("server error: ", err)
		}
	}()

	return server
}

func main() {
	loadEnv()
	connectDB()

	server := startServer()

	// Handle graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	log.Println("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("server shutdown error: ", err)
	}

	log.Println("server gracefully stopped")
}
