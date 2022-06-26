package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "nics",
		Price: 5.5,
		SKU:   "asdf-asdf-xcv",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
