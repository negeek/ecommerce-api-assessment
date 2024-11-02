package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/negeek/ecommerce-api-assessment/repositories"
	"github.com/negeek/ecommerce-api-assessment/services"
	"github.com/negeek/ecommerce-api-assessment/utils"
)

type ProductController struct {
	productService *services.ProductService
}

func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (pc *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	var product repositories.Product
	err := utils.Unmarshall(r.Body, &product)
	if err != nil {
		utils.JsonResponse(w, false, http.StatusBadRequest, "error reading request data", nil)
		return
	}

	err = pc.productService.Create(&product)
	if err != nil {
		utils.JsonResponse(w, false, http.StatusBadGateway, err.Error(), nil)
		return
	}

	utils.JsonResponse(w, true, http.StatusCreated, "Product created successfully", product)
}

func (pc *ProductController) Find(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JsonResponse(w, false, http.StatusBadRequest, "invalid product ID", nil)
		return
	}

	product, err := pc.productService.FindByID(id)
	if err != nil {
		utils.JsonResponse(w, false, http.StatusNotFound, err.Error(), nil)
		return
	}

	utils.JsonResponse(w, true, http.StatusOK, "Product retrieved successfully", product)
}

func (pc *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JsonResponse(w, false, http.StatusBadRequest, "invalid product ID", nil)
		return
	}

	var product repositories.Product
	err = utils.Unmarshall(r.Body, &product)
	if err != nil {
		utils.JsonResponse(w, false, http.StatusBadRequest, "error reading request data", nil)
		return
	}
	product.ID = id

	err = pc.productService.Update(&product, true)
	if err != nil {
		utils.JsonResponse(w, false, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.JsonResponse(w, true, http.StatusOK, "Product updated successfully", nil)
}

func (pc *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JsonResponse(w, false, http.StatusBadRequest, "invalid product ID", nil)
		return
	}

	err = pc.productService.Delete(id)
	if err != nil {
		utils.JsonResponse(w, false, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.JsonResponse(w, true, http.StatusOK, "Product deleted successfully", nil)
}

func (pc *ProductController) List(w http.ResponseWriter, r *http.Request) {
	products, err := pc.productService.List()
	if err != nil {
		utils.JsonResponse(w, false, http.StatusBadGateway, err.Error(), nil)
		return
	}

	utils.JsonResponse(w, true, http.StatusOK, "Products retrieved successfully", products)
}
