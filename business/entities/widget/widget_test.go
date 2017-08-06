package widget_test

import (
	"github.com/danilojunS/widgets-spa-api/business/entities/widget"
	"testing"
)

// Fixtures for user data
const widgetFixtureID = 1
const widgetFixtureName = "My Widget"
const widgetFixtureColor = "blue"
const widgetFixturePrice = "9.99"
const widgetFixtureInventory = 42
const widgetFixtureMelts = true

func TestCreateWidget(t *testing.T) {
	w := widget.Widget{
		ID:        widgetFixtureID,
		Name:      widgetFixtureName,
		Color:     widgetFixtureColor,
		Price:     widgetFixturePrice,
		Inventory: widgetFixtureInventory,
		Melts:     widgetFixtureMelts,
	}
	if err := w.Validate(); err != nil {
		t.Error("Should create a valid widget with ID, Name, Color, Price, Inventory and Melts")
	}
}

func TestCreateWidgetFailMissingId(t *testing.T) {
	w := widget.Widget{
		Name:  widgetFixtureName,
		Color: widgetFixtureColor,
		Price: widgetFixturePrice,
	}
	if err := w.Validate(); err == nil {
		t.Error("Should not create a valid widget without ID")
	}
}

func TestCreateWidgetFailMissingName(t *testing.T) {
	w := widget.Widget{
		ID:    widgetFixtureID,
		Color: widgetFixtureColor,
		Price: widgetFixturePrice,
	}
	if err := w.Validate(); err == nil {
		t.Error("Should not create a valid widget without Name")
	}
}

func TestCreateWidgetFailMissingColor(t *testing.T) {
	w := widget.Widget{
		ID:    widgetFixtureID,
		Name:  widgetFixtureName,
		Price: widgetFixturePrice,
	}
	if err := w.Validate(); err == nil {
		t.Error("Should not create a valid widget without Color")
	}
}

func TestCreateWidgetFailMissingPrice(t *testing.T) {
	w := widget.Widget{
		ID:    widgetFixtureID,
		Name:  widgetFixtureName,
		Color: widgetFixtureColor,
	}
	if err := w.Validate(); err == nil {
		t.Error("Should not create a valid widget without Price")
	}
}
