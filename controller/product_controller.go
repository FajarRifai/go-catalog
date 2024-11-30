package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-catalog/bean"
	"go-catalog/models"
	"go-catalog/service"
	"net/http"
	"strconv"
	"strings"
)

type ProductController struct {
	Service *services.ProductService
}

func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		bean.ErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	id, err := c.Service.CreateProduct(product)
	if err != nil {
		bean.ErrorResponse(w, http.StatusInternalServerError, "Failed to create product")
		return
	}

	bean.JsonResponse(w, http.StatusCreated, "00", "Product created successfully", map[string]int64{"id": id})
}

func (c *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := c.Service.GetProducts()
	if err != nil {
		bean.ErrorResponse(w, http.StatusInternalServerError, "Failed to fetch products")
		return
	}

	bean.JsonResponse(w, http.StatusOK, "00", "Success", products)
}

func (c *ProductController) GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Validate ID
	id, err := strconv.Atoi(idStr)
	if err != nil {
		bean.ErrorResponse(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	product, err := c.Service.GetProductById(id)
	if err != nil {
		bean.ErrorResponse(w, http.StatusNotFound, "Product not found")
		return
	}

	bean.JsonResponse(w, http.StatusOK, "00", "Success", product)
}

func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		bean.ErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	if err := c.Service.UpdateProduct(product); err != nil {
		bean.ErrorResponse(w, http.StatusInternalServerError, "Failed to update product")
		return
	}

	bean.JsonResponse(w, http.StatusOK, "00", "Product updated successfully", nil)
}

func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Validate ID
	id, err := strconv.Atoi(idStr)
	if err != nil {
		bean.ErrorResponse(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := c.Service.DeleteProduct(id); err != nil {
		bean.ErrorResponse(w, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	bean.JsonResponse(w, http.StatusOK, "00", "Product deleted successfully", nil)
}

func (c *ProductController) GetProductByCodes(w http.ResponseWriter, r *http.Request) {
	// Get 'codes' parameter from query string
	codesParam := r.URL.Query().Get("codes")
	fmt.Println("Codes received:", codesParam)
	if codesParam == "" {
		bean.ErrorResponse(w, http.StatusBadRequest, "Product codes are required")
		return
	}

	// Split codes by comma
	codes := strings.Split(codesParam, ",")
	products, err := c.Service.GetProductByCodes(codes)
	if err != nil {
		bean.ErrorResponse(w, http.StatusInternalServerError, "Failed to fetch products by codes")
		return
	}

	if len(products) == 0 {
		bean.JsonResponse(w, http.StatusOK, "00", "No products found for the given codes", nil)
		return
	}

	bean.JsonResponse(w, http.StatusOK, "00", "Success", products)
}
