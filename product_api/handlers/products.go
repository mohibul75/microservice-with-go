package handlers

import (
	"github.com/gorilla/mux"
	"github.com/mohibul75/microservice-with-go/data"
	"log"
	"net/http"
	"strconv"
)

type GetProduct struct {
	l *log.Logger
}

func NewGetProduct(l *log.Logger) *GetProduct {
	return &GetProduct{l}
}

func (p *GetProduct) GetProucts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("get product api ")
	products := data.GetProducts()
	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal product", http.StatusBadRequest)
	}
}

func (p *GetProduct) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("add product added")
	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to marshal", http.StatusBadRequest)
	}

	data.AddProduct(product)
}

func (p *GetProduct) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Product id", http.StatusBadRequest)
	}

	p.l.Println("Update product api")
	product := &data.Product{}
	err = product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Update product Failed", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, product)
	if err != nil {
		http.Error(w, "Update Product Failed", http.StatusBadRequest)
	}
}
