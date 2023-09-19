package main

import (
	"context"
	"github.com/mohibul75/microservice-with-go/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	l := log.New(os.Stdout, "product-api	:	", log.LstdFlags)
	surveMux := http.NewServeMux()

	homeHandler := handlers.NewHome(l)
	dashboardHanlder := handlers.NewDashboard(l)
	productsHandler := handlers.NewGetProduct(l)

	surveMux.Handle("/", homeHandler)
	surveMux.Handle("/dashboard", dashboardHanlder)
	surveMux.Handle("/getproducts", productsHandler)

	server := http.Server{
		Addr:         ":4000",
		Handler:      surveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		l.Println("Server started at port 4000")

		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	sig := <-sigChannel
	l.Print("Shutdown is called gracefully	", sig)

	tx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tx)
}
