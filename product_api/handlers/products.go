package handlers

import (
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

func (p *GetProduct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProucts(w, r)
		return
	}

	w.WriteHeader(http.StatusNotImplemented)
}

func (p *GetProduct) getProucts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("get product api ")
	products := data.GetProducts()
	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal product", http.StatusBadRequest)
	}
}
