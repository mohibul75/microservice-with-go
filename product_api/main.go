package main

import (
	"github.com/mohibul75/microservice-with-go/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	surveMux := http.NewServeMux()

	homeHandler := handlers.NewHome(l)
	dashboardHanlder := handlers.NewDashboard(l)

	surveMux.Handle("/", homeHandler)
	surveMux.Handle("/dashboard", dashboardHanlder)

	server := http.Server{
		Addr:         ":4000",
		Handler:      surveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	server.ListenAndServe()
}
