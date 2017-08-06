package useCase

import (
	widget "github.com/danilojunS/widgets-spa-api/business/entities/widget"
	widgetRepo "github.com/danilojunS/widgets-spa-api/business/repositories/widget"
)

// GetWidgets without filtering use case
func GetWidgets() ([]widget.Widget, error) {
	widgets, err := widgetRepo.Read(0)

	if err != nil {
		return nil, err
	}

	return widgets, nil
}
