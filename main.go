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
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Oops"))
			return
		}
		fmt.Fprint(w, "Hello %s", data)

	})

	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Bye")
	})

	http.ListenAndServe(":4000", nil)
}
