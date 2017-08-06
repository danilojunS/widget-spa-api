package useCase

import (
	"errors"
	widget "github.com/danilojunS/widgets-spa-api/business/entities/widget"
	widgetRepo "github.com/danilojunS/widgets-spa-api/business/repositories/widget"
)

// UpdateWidget use case
func UpdateWidget(
	id int,
	name string,
	color string,
	price string,
	inventory int,
	melts bool,
) (widget.Widget, error) {

	widgets, err := widgetRepo.Read(id)
	if err != nil {
		return widget.Widget{}, errors.New("Failed to find widget in database")
	}

	if len(widgets) == 0 {
		return widget.Widget{}, errors.New("Widget ID not found")
	}

	w := widgets[0]

	w.Name = name
	w.Color = color
	w.Price = price
	w.Inventory = inventory
	w.Melts = melts

	_, err = widgetRepo.Update(w)
	if err != nil {
		return widget.Widget{}, errors.New("Failed to update widget in database")
	}

	return w, nil
}
