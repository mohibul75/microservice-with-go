package handlers

import (
	"github.com/gorilla/mux"
	"github.com/mohibul75/microservice-with-go/data"
	"net/http"
	"strconv"
)

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
