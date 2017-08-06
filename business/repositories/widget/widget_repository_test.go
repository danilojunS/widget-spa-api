package widgetRepo_test

import (
	widget "github.com/danilojunS/widgets-spa-api/business/entities/widget"
	widgetRepo "github.com/danilojunS/widgets-spa-api/business/repositories/widget"
	utils "github.com/danilojunS/widgets-spa-api/utils"
	"testing"
)

// Fixtures for widget data
const (
	widgetFixtureName      = "My Widget"
	widgetFixtureColor     = "blue"
	widgetFixturePrice     = "9.99"
	widgetFixtureInventory = 42
	widgetFixtureMelts     = true
)

var createdWidgetID int
var createdWidget widget.Widget
var updatedWidget widget.Widget

func TestWidgetCreate(t *testing.T) {
	u := widget.Widget{
		Name:      widgetFixtureName,
		Color:     widgetFixtureColor,
		Price:     widgetFixturePrice,
		Inventory: widgetFixtureInventory,
		Melts:     widgetFixtureMelts,
	}
	var err error
	createdWidgetID, err = widgetRepo.Create(u)
	utils.CheckError(err)
}

func TestWidgetRead(t *testing.T) {
	widgets, err := widgetRepo.Read(createdWidgetID)
	utils.CheckError(err)

	if len(widgets) != 1 {
		t.Error("Should find one created widget")
		return
	}

	createdWidget = widgets[0]

	if createdWidget.ID != createdWidgetID ||
		createdWidget.Name != widgetFixtureName ||
		createdWidget.Color != widgetFixtureColor ||
		createdWidget.Price != widgetFixturePrice ||
		createdWidget.Inventory != widgetFixtureInventory ||
		createdWidget.Melts != widgetFixtureMelts {
		t.Error("Should find created widget")
	}
}

func TestWidgetUpdate(t *testing.T) {
	const newColor string = "red"
	createdWidget.Color = newColor
	updatedWidgetID, err := widgetRepo.Update(createdWidget)
	utils.CheckError(err)

	if updatedWidgetID != createdWidget.ID {
		t.Error("Should update correct widget")
		return
	}

	widgets, err := widgetRepo.Read(createdWidget.ID)
	utils.CheckError(err)

	if len(widgets) != 1 {
		t.Error("Should find one updated widget")
		return
	}

	updatedWidget = widgets[0]

	if updatedWidget.ID != createdWidgetID || updatedWidget.Color != newColor {
		t.Error("Should update widget with correct attributes")
	}
}

func TestDeleteWidget(t *testing.T) {
	deletedWidgetID, err := widgetRepo.Delete(updatedWidget)
	utils.CheckError(err)

	if deletedWidgetID != updatedWidget.ID {
		t.Error("Should delete correct widget")
		return
	}

	widgets, err := widgetRepo.Read(deletedWidgetID)
	utils.CheckError(err)

	if len(widgets) > 0 {
		t.Error("Should not find deleted widget")
	}
}

func TestClearWidgets(t *testing.T) {
	widgets := []widget.Widget{
		{Name: "widget1", Color: "color1", Price: "1.99"},
		{Name: "widget2", Color: "color2", Price: "2.99"},
		{Name: "widget3", Color: "color3", Price: "3.99"},
	}

	for _, widget := range widgets {
		_, err := widgetRepo.Create(widget)
		utils.CheckError(err)
	}

	err := widgetRepo.Clear()
	utils.CheckError(err)

	widgetsFromRepo, err := widgetRepo.Read(0)
	utils.CheckError(err)

	if len(widgetsFromRepo) > 0 {
		t.Error("Should clear all widgets in repository")
	}
}

func TestTeardown(t *testing.T) {
	err := widgetRepo.Clear()
	utils.CheckError(err)
}
