package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/justiniso/golock/reservation"
	"github.com/justiniso/golock/server"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Request: %s!", r.URL.Path[1:])
}

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter()
	r.HandleFunc("/reservation/{resource}", reservation.ReservationHandler).
		Methods("GET", "POST")

	server := server.NewServer(host, port, r)

	log.Printf("Server listening on: %s", server.Address())
	server.Start()
	log.Fatal("Server died...")
}
