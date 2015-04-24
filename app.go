package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/justiniso/golock/urls"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Request: %s!", r.URL.Path[1:])
}

func main() {
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	urls.Handle()
}
