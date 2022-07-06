package data

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
	"regexp"
)

//region struct

// Product defines the structure for an API product
// https://pkg.go.dev/encoding/json#Marshal definition for JSON Marshal
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return validate.Struct(p)
}

//endregion struct

//region utilities
func validateSKU(fieldLevel validator.FieldLevel) bool {
	format := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := format.FindAllString(fieldLevel.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}
	return true
}

//endregion utilities
