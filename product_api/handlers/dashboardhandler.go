package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Dashboard struct {
	l *log.Logger
}

func NewDashboard(l *log.Logger) *Dashboard {
	return &Dashboard{l}
}

func (dashboard *Dashboard) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dashboard.l.Println("Welcome to dashboard")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		dashboard.l.Println("Error in request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	dashboard.l.Printf("data: %s", data)
	fmt.Fprintf(w, "Data: %s", data)
}
