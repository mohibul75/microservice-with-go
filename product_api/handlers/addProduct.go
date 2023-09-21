package handlers

import (
	"github.com/mohibul75/microservice-with-go/data"
	"net/http"
)

func (p *Product) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("add product api")
	product := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&product)
}
