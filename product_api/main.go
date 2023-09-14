package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Home")

		data, err := ioutil.ReadAll(r.Body)
		log.Printf("Data %s", data)

		if err != nil {
			// w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("Oops"))
			http.Error(w, "Ooopss", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello %s\n", data)

	})

	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Bye")
	})

	http.ListenAndServe(":4000", nil)
}
