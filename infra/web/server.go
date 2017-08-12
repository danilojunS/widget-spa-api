package web

import (
	"fmt"
	config "github.com/danilojunS/widgets-spa-api/config"
	handlers "github.com/danilojunS/widgets-spa-api/infra/web/handlers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
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
		// bypass authorization if config does not require it
		if !config.Get().Auth {
			handler(w, req)
			return
		}

		// authorization header is required
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			handlers.UnauthorizedError(w, "")
			return
		}

		// user is not sending the token type and token
		// in the format: "Authorization: <type> <token>"
		tokenWords := strings.Fields(authHeader)
		if len(tokenWords) != 2 {
			handlers.UnauthorizedError(w, "")
			return
		}

		// only supported token type is "bearer"
		tokenType := tokenWords[0]
		if strings.ToLower(tokenType) != "bearer" {
			handlers.UnauthorizedError(w, "")
			return
		}

		// verify token
		tokenString := tokenWords[1]
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

			handlers.UnauthorizedError(w, "")
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
