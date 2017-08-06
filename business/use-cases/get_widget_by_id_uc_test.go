package useCase_test

import (
	widget "github.com/danilojunS/widgets-spa-api/business/entities/widget"
	widgetRepo "github.com/danilojunS/widgets-spa-api/business/repositories/widget"
	useCases "github.com/danilojunS/widgets-spa-api/business/use-cases"
	utils "github.com/danilojunS/widgets-spa-api/utils"
	"testing"
)

// Get Widgets by ID
func TestGetWidgetByID(t *testing.T) {
	defer func() {
		err := widgetRepo.Clear()
		utils.CheckError(err)
	}()

	const (
		widgetName      = "My Widget"
		widgetColor     = "blue"
		widgetPrice     = "9.99"
		widgetInventory = 42
		widgetMelts     = true
	)

	// create dummy widget
	widget := widget.Widget{
		Name:      widgetName,
		Color:     widgetColor,
		Price:     widgetPrice,
		Inventory: widgetInventory,
		Melts:     widgetMelts,
	}
	id, err := widgetRepo.Create(widget)
	utils.CheckError(err)

	widgetFromUseCase, err := useCases.GetWidgetByID(id)
	utils.CheckError(err)

	if widgetFromUseCase.ID != id ||
		widgetFromUseCase.Name != widgetName ||
		widgetFromUseCase.Color != widgetColor ||
		widgetFromUseCase.Price != widgetPrice ||
		widgetFromUseCase.Inventory != widgetInventory ||
		widgetFromUseCase.Melts != widgetMelts {
		t.Error("Should get created widget by ID in use case")
	}
}
