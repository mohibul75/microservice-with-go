package handlers

import (
	"github.com/mohibul75/microservice-with-go/data"
	"net/http"
)

func (p *Product) GetProucts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("get product api ")
	products := data.GetProducts()
	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal product", http.StatusBadRequest)
		return
	}
}
