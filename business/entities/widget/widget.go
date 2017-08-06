package widget

import "errors"

// Widget business entity
type Widget struct {
	ID        int
	Name      string
	Color     string
	Price     string
	Inventory int
	Melts     bool
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
