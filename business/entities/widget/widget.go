package widget

import "errors"

// Widget business entity
type Widget struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	Price     string `json:"price"`
	Inventory int    `json:"inventory"`
	Melts     bool   `json:"melts"`
}

// Validate a widget
func (w Widget) Validate() error {
	if w.Name == "" {
		return errors.New("name is a required field and must not be empty string")
	}
	if w.Color == "" {
		return errors.New("color is a required field and must not be empty string")
	}
	if w.Price == "" {
		return errors.New("price is a required field and must not be empty string")
	}

	return nil
}
