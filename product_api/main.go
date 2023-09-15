package main

import (
	"github.com/mohibul75/microservice-with-go/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	surveMux := http.NewServeMux()

	homeHandler := handlers.NewHome(l)
	dashboardHanlder := handlers.NewDashboard(l)

	surveMux.Handle("/", homeHandler)
	surveMux.Handle("/dashboard", dashboardHanlder)

	http.ListenAndServe(":4000", surveMux)
}
