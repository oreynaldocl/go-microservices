package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

var ErrProductNotFound = fmt.Errorf("product not found")

//region struct Products

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	p.privateExample("Just example")
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p Products) privateExample(message string) {
	fmt.Println("print message", message)
}

//endregion struct Products

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, prod *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	prod.ID = id
	productList[pos] = prod
	return nil
}

func findProduct(id int) (*Product, int, error) {
	for index, prod := range productList {
		if prod.ID == id {
			return prod, index, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// productList hardcoded example for simplify example
var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
