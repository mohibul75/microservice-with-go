package main

import (
	"github.com/mohibul75/microservice-with-go/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	homeHandler := handlers.NewHome(l)
	surveMux := http.NewServeMux()
	surveMux.Handle("/", homeHandler)

	//http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
	//	log.Print("Bye")
	//})

	http.ListenAndServe(":4000", surveMux)
}
