package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Home struct {
	l *log.Logger
}

func NewHome(l *log.Logger) *Home {
	return &Home{l}
}

func (home *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	home.l.Print("Home")

	data, err := ioutil.ReadAll(r.Body)
	home.l.Printf("Recieved Data %s", data)

	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write([]byte("Oops"))
		home.l.Print("Error Reading Request Data")
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello %s\n", data)
}
