package web

import (
	handlers "github.com/danilojunS/widgets-spa-api/infra/web/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

// StartServer starts web server
func StartServer() error {
	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.UserGet).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UserGet).Methods("GET")

	return http.ListenAndServe(":4000", r)
}
