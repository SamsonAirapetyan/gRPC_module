package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreateOn    string  `json:"-"`
	UpdateOn    string  `json:"-"`
	DeleteOn    string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProduct() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Milk + Coffee",
		Price:       220,
		SKU:         "avb4",
		CreateOn:    time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Coffee",
		Price:       110,
		SKU:         "asd2",
		CreateOn:    time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
}
