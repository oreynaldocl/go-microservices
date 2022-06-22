package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"working/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (ph *Products) ServeHTTP(rw http.ResponseWriter, rh *http.Request) {
	if rh.Method == http.MethodGet {
		ph.getProducts(rw, rh)
		return
	}

	if rh.Method == http.MethodPost {
		ph.addProduct(rw, rh)
		return
	}

	if rh.Method == http.MethodPut {
		r := regexp.MustCompile(`/(\d+)`)
		g := r.FindAllStringSubmatch(rh.URL.Path, -1)
		if len(g) != 1 || len(g[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		ph.updateProduct(id, rw, rh)
		return
	}

	// catch not implemented
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (ph *Products) getProducts(rw http.ResponseWriter, _ *http.Request) {
	ph.l.Println("Handle GET Products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (ph *Products) addProduct(rw http.ResponseWriter, rh *http.Request) {
	ph.l.Println("Handle POST Product")

	product := &data.Product{}
	err := product.FromJSON(rh.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	data.AddProduct(product)
}

func (ph *Products) updateProduct(id int, rw http.ResponseWriter, rh *http.Request) {
	ph.l.Println("Handle PUT Product")

	product := &data.Product{}
	err := product.FromJSON(rh.Body)
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
