package handlers

import (
	"encoding/json"
	"github.com/mohibul75/microservice-with-go/data"
	"log"
	"net/http"
)

type GetProduct struct {
	l *log.Logger
}

func NewGetProduct(l *log.Logger) *GetProduct {
	return &GetProduct{l}
}

func (getProduct *GetProduct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	getProduct.l.Println("get product api ")
	products := data.GetProducts()
	data, err := json.Marshal(products)
	if err != nil {
		http.Error(w, "Unable to marshal product", http.StatusBadRequest)
	}

	w.Write(data)
}
