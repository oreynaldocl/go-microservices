package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"working/product-api/data"
)

type Products struct {
	log *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (ph *Products) GetProducts(rw http.ResponseWriter, _ *http.Request) {
	ph.log.Println("Handle GET Products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (ph *Products) AddProduct(rw http.ResponseWriter, rh *http.Request) {
	ph.log.Println("Handle POST Product")

	product := &data.Product{}
	err := product.FromJSON(rh.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	data.AddProduct(product)
}

func (ph *Products) UpdateProduct(rw http.ResponseWriter, rh *http.Request) {
	vars := mux.Vars(rh)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, fmt.Sprintf("Unable to convert id=%s", vars["id"]), http.StatusBadRequest)
		return
	}

	ph.log.Println("Handle PUT Product", id)

	product := &data.Product{}
	err = product.FromJSON(rh.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, product)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Error while PUT", http.StatusInternalServerError)
		return
	}

}
