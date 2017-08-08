package widgetRepo

import (
	"github.com/danilojunS/widgets-spa-api/business/entities/widget"
)

var records []widget.Widget
var lastInsertedID = 1

// Create a widget
func CreateMock(w widget.Widget) (int, error) {
	w.ID = lastInsertedID
	lastInsertedID++
	records = append(records, w)
	return w.ID, nil
}

// Read widgets
func ReadMock(id int) ([]widget.Widget, error) {
	if id == 0 {
		return records, nil
	}

	var matches []widget.Widget
	for _, r := range records {
		if r.ID == id {
			matches = append(matches, r)
		}
	}

	return matches, nil
}

// Update widget
func UpdateMock(w widget.Widget) (int, error) {
	for i, r := range records {
		if r.ID == w.ID {
			records[i] = w
		}
	}
	return w.ID, nil
}

// Delete widget
func DeleteMock(w widget.Widget) (int, error) {
	for i, r := range records {
		if r.ID == w.ID {
			records = append(records[:i], records[i+1:]...)
		}
	}
	return w.ID, nil
}

// Clear all widgets
func ClearMock() error {
	records = []widget.Widget{}
	return nil
}
