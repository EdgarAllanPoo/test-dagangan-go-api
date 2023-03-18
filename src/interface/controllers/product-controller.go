package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/EdgarAllanPoo/test-go-api/src/domain"

	"github.com/EdgarAllanPoo/test-go-api/src/usecases"
)

type ProductController struct {
	productInteractor usecases.ProductInteractor
}

func NewProductController(productInteractor usecases.ProductInteractor) *ProductController {
	return &ProductController{productInteractor}
}

func (controller *ProductController) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var product domain.Product
	err := json.NewDecoder(req.Body).Decode(&product)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}
	err2 := controller.productInteractor.CreateProduct(product)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (controller *ProductController) FindAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	results, err2 := controller.productInteractor.FindAll()
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}

func (controller *ProductController) FindById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	idStr := vars["id"]

	// Parse the string ID to an int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid product ID"})
		return
	}

	product, err := controller.productInteractor.FindById(id)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(product)
}

func (controller *ProductController) Delete(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	idStr := vars["id"]

	// Parse the string ID to an int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid product ID"})
		return
	}

	err2 := controller.productInteractor.DeleteProduct(id)
	if err2 != nil {
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (controller *ProductController) Update(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	idStr := vars["id"]

	// Parse the string ID to an int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid product ID"})
		return
	}

	var updatedProduct domain.Product
	err = json.NewDecoder(req.Body).Decode(&updatedProduct)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}

	updatedProduct.Id = id
	err2 := controller.productInteractor.UpdateProduct(id, updatedProduct)
	if err2 != nil {
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (controller *ProductController) FilterByCategory(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	category := req.URL.Query().Get("category")
	if category == "" {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Missing category parameter"})
		return
	}

	results, err := controller.productInteractor.FilterByCategory(category)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}
