package web

import (
	"fmt"
	handlers "github.com/danilojunS/widgets-spa-api/infra/web/handlers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
)

// StartServer starts web server
func StartServer() error {
	r := mux.NewRouter()

	r.HandleFunc("/users", secure(handlers.UserGet)).Methods("GET")
	r.HandleFunc("/users/{id}", secure(handlers.UserGetByID)).Methods("GET")

	r.HandleFunc("/widgets", secure(handlers.WidgetGet)).Methods("GET")
	r.HandleFunc("/widgets/{id}", secure(handlers.WidgetGetByID)).Methods("GET")
	r.HandleFunc("/widgets", secure(handlers.WidgetPost)).Methods("POST")
	r.HandleFunc("/widgets/{id}", secure(handlers.WidgetPut)).Methods("PUT")

	r.HandleFunc("/token", handlers.TokenGet).Methods("GET")

	return http.ListenAndServe(":4000", r)
}

// middleware for token validation
func secure(handler func(w http.ResponseWriter, req *http.Request)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		tokenString := req.Header.Get("Authorization")
		if tokenString == "" {
			handlers.UnauthorizedError(w, "")
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(handlers.TokenSecret), nil
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

		claims, ok := token.Claims.(jwt.MapClaims)
		fmt.Println(claims["foo"], claims["nbf"])

		if !ok || !token.Valid {
			handlers.UnauthorizedError(w, "")
			return
		}

		handler(w, req)
	}
}
