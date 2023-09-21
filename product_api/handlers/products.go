package handlers

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mohibul75/microservice-with-go/data"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) GetProucts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("get product api ")
	products := data.GetProducts()
	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal product", http.StatusBadRequest)
		return
	}
}

func (p *Product) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("add product api")
	product := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&product)
}

func (p *Product) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Product id", http.StatusBadRequest)
		return
	}

	p.l.Println("Update product api")
	product := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &product)
	if err != nil {
		http.Error(w, "Update Product Failed", http.StatusBadRequest)
		return
	}
}

type KeyProduct struct{}

func (p *Product) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR]	deserializing product", err)
			http.Error(w, "Error on reading product ", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] Product validation Failed", err)
			http.Error(
				w,
				fmt.Sprintf("Product Validation Falied : %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
