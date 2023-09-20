package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeleteOn    string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(writer io.Writer) error {
	e := json.NewEncoder(writer)
	return e.Encode(p)
}

func (p *Product) FromJSON(reader io.Reader) error {
	e := json.NewDecoder(reader)
	return e.Decode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = findNextID()
	productList = append(productList, p)
}

func findNextID() int {
	lastProduct := productList[len(productList)-1]
	return lastProduct.ID + 1
}

func UpdateProduct(id int, p *Product) error {
	prod, pos, err := findId(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = prod

	return nil
}

var ProductNotFoundError = fmt.Errorf("Product Not Found")

func findId(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ProductNotFoundError
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
