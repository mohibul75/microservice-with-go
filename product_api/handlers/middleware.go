package handlers

import (
	"context"
	"fmt"
	"github.com/mohibul75/microservice-with-go/data"
	"net/http"
)

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
