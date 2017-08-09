package web

import (
	"fmt"
	config "github.com/danilojunS/widgets-spa-api/config"
	handlers "github.com/danilojunS/widgets-spa-api/infra/web/handlers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// StartServer starts web server
func StartServer(port int) error {
	r := mux.NewRouter()

	r.HandleFunc("/users", secure(handlers.UserGet)).Methods("GET")
	r.HandleFunc("/users/{id}", secure(handlers.UserGetByID)).Methods("GET")

	r.HandleFunc("/widgets", secure(handlers.WidgetGet)).Methods("GET")
	r.HandleFunc("/widgets/{id}", secure(handlers.WidgetGetByID)).Methods("GET")
	r.HandleFunc("/widgets", secure(handlers.WidgetPost)).Methods("POST")
	r.HandleFunc("/widgets/{id}", secure(handlers.WidgetPut)).Methods("PUT")

	r.HandleFunc("/token", handlers.TokenGet).Methods("GET")

	return http.ListenAndServe(fmt.Sprint(":", strconv.Itoa(port)), r)
}

// middleware for token validation
func secure(handler func(w http.ResponseWriter, req *http.Request)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		if !config.Get().Auth {
			handler(w, req)
			return
		}

		tokenString := req.Header.Get("Authorization")
		if tokenString == "" {
			handlers.UnauthorizedError(w, "")
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(config.Get().TokenSecret), nil
		})
		if err != nil {
			errMessage := fmt.Sprint(err)
			if errMessage == "Token is expired" {
				handlers.UnauthorizedError(w, errMessage)
				return
			}

			handlers.InternalError(w, "")
			return
		}

		_, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			handlers.UnauthorizedError(w, "")
			return
		}

		handler(w, req)
	}
}
