package handlers

import (
	"encoding/json"
	useCases "github.com/danilojunS/widgets-spa-api/business/use-cases"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// UserGet handler
func UserGet(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	id := params["id"]

	// get user by id
	if id != "" {
		intID, _ := strconv.Atoi(id)
		user, _ := useCases.GetUserByID(intID)
		log.Println(user)
		json.NewEncoder(w).Encode(user)
		return
	}

	users, _ := useCases.GetUsers()

	if len(users) == 0 {
		json.NewEncoder(w).Encode([]string{})
		return
	}

	json.NewEncoder(w).Encode(users)
}
