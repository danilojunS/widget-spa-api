package useCase

import (
	"errors"
	widget "github.com/danilojunS/widgets-spa-api/business/entities/widget"
	widgetRepo "github.com/danilojunS/widgets-spa-api/business/repositories/widget"
)

// CreateWidget use case
func CreateWidget(
	name string,
	color string,
	price string,
	inventory int,
	melts bool,
) (widget.Widget, error) {
	w := widget.Widget{
		Name:      name,
		Color:     color,
		Price:     price,
		Inventory: inventory,
		Melts:     melts,
	}

	err := w.Validate()
	if err != nil {
		return widget.Widget{}, errors.New("Invalid parameters")
	}

	id, err := widgetRepo.Create(w)
	if err != nil {
		return widget.Widget{}, err
	}

	w.ID = id
	return w, nil
}
