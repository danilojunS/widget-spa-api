package widgetRepo

import (
	widget "github.com/danilojunS/widgets-spa-api/business/entities/widget"
	config "github.com/danilojunS/widgets-spa-api/config"
)

// Create a widget
func Create(w widget.Widget) (int, error) {
	if config.Get().DBMock {
		return CreateMock(w)
	}
	return CreatePg(w)
}

// Read widgets
func Read(id int) ([]widget.Widget, error) {
	if config.Get().DBMock {
		return ReadMock(id)
	}
	return ReadPg(id)
}

// Update widget
func Update(w widget.Widget) (int, error) {
	if config.Get().DBMock {
		return UpdateMock(w)
	}
	return UpdatePg(w)
}

// Delete widget
func Delete(w widget.Widget) (int, error) {
	if config.Get().DBMock {
		return DeleteMock(w)
	}
	return DeletePg(w)
}

// Clear all widgets
func Clear() error {
	if config.Get().DBMock {
		return ClearMock()
	}
	return ClearPg()
}
