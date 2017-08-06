package useCase

import (
	"errors"
	widget "github.com/danilojunS/widgets-spa-api/business/entities/widget"
	widgetRepo "github.com/danilojunS/widgets-spa-api/business/repositories/widget"
)

// GetWidgetByID use case
func GetWidgetByID(id int) (widget.Widget, error) {
	widgets, err := widgetRepo.Read(id)

	if err != nil {
		return widget.Widget{}, err
	}

	if len(widgets) == 0 {
		return widget.Widget{}, errors.New("No widget found for ID")
	}

	return widgets[0], nil
}
