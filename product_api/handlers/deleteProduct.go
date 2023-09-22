package handlers

import (
	"github.com/mohibul75/microservice-with-go/data"
	"net/http"
)

func (p *Product) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	p.l.Printf("[API] : handle delete product api and product id : %s", id)
	err := data.RemoveProduct(id)
	if err == data.ProductNotFoundError {
		p.l.Println("[ERROR] : product not found")
		http.Error(w, "invalid product id", http.StatusNotFound)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] : product delete failed")
		http.Error(w, "[ERROR] : product delete failed", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
