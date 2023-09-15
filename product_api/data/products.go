package data

import (
	"time"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedOn   string
	UpdatedOn   string
	DeleteOn    string
}

var productList = []*Product{

	&Product{
		ID:          1,
		Name:        "Capachino",
		Description: "Milk Coffee",
		Price:       2.45,
		SKU:         "dcnjd",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Capachino expresso",
		Description: "premium Coffee",
		Price:       6.45,
		SKU:         "dfbu",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
