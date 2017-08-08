package handlers

import (
	"encoding/json"
	useCases "github.com/danilojunS/widgets-spa-api/business/use-cases"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// WidgetGet handler
func WidgetGet(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	widgets, err := useCases.GetWidgets()
	if err != nil {
		InternalError(w, "")
		return
	}

	if len(widgets) == 0 {
		err = json.NewEncoder(w).Encode([]string{})
		if err != nil {
			InternalError(w, "")
		}
		return
	}

	err = json.NewEncoder(w).Encode(widgets)
	if err != nil {
		InternalError(w, "")
	}
}

// WidgetGetByID handler
func WidgetGetByID(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	id := params["id"]

	intID, err := strconv.Atoi(id)
	if err != nil {
		ValidationError(w, ".id must be an int")
		return
	}

	widget, err := useCases.GetWidgetByID(intID)
	if err != nil {
		NotFoundError(w, "")
		return
	}

	err = json.NewEncoder(w).Encode(widget)
	if err != nil {
		InternalError(w, "")
	}
}

// WidgetPost handler
func WidgetPost(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	var wr WidgetRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&wr); err != nil {
		ValidationError(w, "invalid body payload")
		return
	}
	defer req.Body.Close()

	_, err := useCases.CreateWidget(wr.Name, wr.Color, wr.Price, wr.Inventory, wr.Melts)
	if err != nil {
		InternalError(w, "")
		return
	}

	_, err = w.Write([]byte("added!"))
	if err != nil {
		panic(err)
	}
}

// WidgetRequest request for widget creation/update
type WidgetRequest struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	Price     string `json:"price"`
	Inventory int    `json:"inventory"`
	Melts     bool   `json:"melts"`
}
