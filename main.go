package main

import (
	"go-catalog/controller"
	"go-catalog/database"
	"go-catalog/repository"
	"go-catalog/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize database
	database.ConnectDatabase()

	// Initialize layers
	repo := &repository.ProductRepository{DB: database.DB}
	service := &services.ProductService{Repo: repo}
	restController := &controller.ProductController{Service: service}

	// Setup routes
	router := mux.NewRouter()
	router.HandleFunc("/api/product", restController.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products", restController.GetAllProducts).Methods("GET")
	router.HandleFunc("/api/product/{id}", restController.GetProductById).Methods("GET")
	router.HandleFunc("/api/product", restController.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/product-codes", restController.GetProductByCodes).Methods("GET")
	router.HandleFunc("/api/product/{id}", restController.DeleteProduct).Methods("DELETE")

	// Start server
	log.Println("Server running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
