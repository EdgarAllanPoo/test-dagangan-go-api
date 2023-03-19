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
	err = controller.productInteractor.CreateProduct(product)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(SuccessResponse{Message: "Successfully insert new object into product"})
}

func (controller *ProductController) FindAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	category := req.URL.Query().Get("category")
	limitStr := req.URL.Query().Get("limit")
	offsetStr := req.URL.Query().Get("offset")

	var results []*domain.Product
	var totalRows int64
	var err error

	defaultLimit := 10
	defaultOffset := 0

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = defaultLimit
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = defaultOffset
	}

	results, totalRows, err = controller.productInteractor.FindAll(category, limit, offset)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]interface{}{
		"data":       results,
		"total_rows": totalRows,
	})
}

func (controller *ProductController) FindById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	idStr := vars["id"]

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

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid product ID"})
		return
	}

	err = controller.productInteractor.DeleteProduct(id)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(SuccessResponse{Message: "Successfully deleted object with Id " + idStr})
}

func (controller *ProductController) Update(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	idStr := vars["id"]

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
	err = controller.productInteractor.UpdateProduct(id, updatedProduct)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(SuccessResponse{Message: "Successfully updated object with Id " + idStr})

}
