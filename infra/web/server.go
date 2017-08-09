package web

import (
	handlers "github.com/danilojunS/widgets-spa-api/infra/web/handlers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
)

// StartServer starts web server
func StartServer() error {
	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.UserGet).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UserGetByID).Methods("GET")

	r.HandleFunc("/widgets", handlers.WidgetGet).Methods("GET")
	r.HandleFunc("/widgets/{id}", handlers.WidgetGetByID).Methods("GET")
	r.HandleFunc("/widgets", handlers.WidgetPost).Methods("POST")
	r.HandleFunc("/widgets/{id}", handlers.WidgetPut).Methods("PUT")

	r.HandleFunc("/token", handlers.TokenGet).Methods("GET")

	return http.ListenAndServe(":4000", r)
}
