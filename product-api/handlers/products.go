package handlers

import (
	"context"
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

func (ph Products) GetProducts(rw http.ResponseWriter, _ *http.Request) {
	ph.log.Println("Handle GET Products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (ph Products) AddProduct(rw http.ResponseWriter, rh *http.Request) {
	ph.log.Println("Handle POST Product")

	product := rh.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&product)
}

func (ph Products) UpdateProduct(rw http.ResponseWriter, rh *http.Request) {
	vars := mux.Vars(rh)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, fmt.Sprintf("Unable to convert id=%s", vars["id"]), http.StatusBadRequest)
		return
	}

	ph.log.Println("Handle PUT Product", id)

	product := rh.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &product)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Error while PUT", http.StatusInternalServerError)
		return
	}

}

type KeyProduct struct{}

func (ph Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	// ph *Products required for modify the ph class, without it, it is just readonly
	return http.HandlerFunc(func(rw http.ResponseWriter, rh *http.Request) {
		ph.log.Println("Handle MiddlewareValidateProduct")

		// Create a product
		product := data.Product{}
		// Create a pointer of the product
		//product := &data.Product{} // *product := &data.Product{}
		err := product.FromJSON(rh.Body)
		if err != nil {
			ph.log.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		err = product.Validate()
		if err != nil {
			ph.log.Println("[ERROR] validating product", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(rh.Context(), KeyProduct{}, product)
		rh = rh.WithContext(ctx)

		next.ServeHTTP(rw, rh)
	})
}
