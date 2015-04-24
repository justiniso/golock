// Urls is the package for maintaining all url routing for the application
package urls

import (
	"github.com/gorilla/mux"
	"github.com/justiniso/golock/reservation"
	"net/http"
)

func Handle() {
	r := mux.NewRouter()
	r.HandleFunc("/reservation/{resource}", reservation.ReservationHandler).
		Methods("GET", "POST")
	http.handle("/", r)
}
