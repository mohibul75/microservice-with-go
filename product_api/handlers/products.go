package handlers

import (
	"github.com/mohibul75/microservice-with-go/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
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

	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		p.updateProduct(id, w, r)
	}
	p.l.Printf("Method Type : %s ", r.Method)
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

func (p *GetProduct) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("add product added")
	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to marshal", http.StatusBadRequest)
	}

	data.AddProduct(product)
}

func (p *GetProduct) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("Update product api")
	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Update product Failed", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, product)
	if err != nil {
		http.Error(w, "Update Product Failed", http.StatusBadRequest)
	}
}
