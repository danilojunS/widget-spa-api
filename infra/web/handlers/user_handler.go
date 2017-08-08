package handlers

import (
	"encoding/json"
	useCases "github.com/danilojunS/widgets-spa-api/business/use-cases"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// UserGet handler
func UserGet(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := useCases.GetUsers()
	if err != nil {
		InternalError(w, "")
		return
	}

	if len(users) == 0 {
		err = json.NewEncoder(w).Encode([]string{})
		if err != nil {
			InternalError(w, "")
		}
		return
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		InternalError(w, "")
	}
}

// UserGetByID handler
func UserGetByID(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	id := params["id"]

	intID, err := strconv.Atoi(id)
	if err != nil {
		ValidationError(w, ".id must be an int")
		return
	}

	user, err := useCases.GetUserByID(intID)
	if err != nil {
		NotFoundError(w, "")
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		InternalError(w, "")
	}
}
