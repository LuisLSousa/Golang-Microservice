package data

import "testing"

func Test_CheckValidation( t *testing.T) {
	p := &Product{
		Name: "Lee",
		Price: 1.00,
		SKU: "abc-def-ghi",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}