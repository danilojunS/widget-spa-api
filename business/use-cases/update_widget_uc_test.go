package useCase_test

import (
	widget "github.com/danilojunS/widgets-spa-api/business/entities/widget"
	widgetRepo "github.com/danilojunS/widgets-spa-api/business/repositories/widget"
	useCases "github.com/danilojunS/widgets-spa-api/business/use-cases"
	utils "github.com/danilojunS/widgets-spa-api/utils"
	"testing"
)

// TestUpdateWidget
func TestUpdateWidget(t *testing.T) {
	defer func() {
		err := widgetRepo.Clear()
		utils.CheckError(err)
	}()

	const (
		name      = "my widget"
		color     = "blue"
		price     = "9.99"
		inventory = 42
		melts     = true
	)

	id, err := widgetRepo.Create(widget.Widget{
		Name:      name,
		Color:     color,
		Price:     price,
		Inventory: inventory,
		Melts:     melts,
	})
	utils.CheckError(err)

	const (
		newName      = "my new widget"
		newColor     = "red"
		newPrice     = "19.99"
		newInventory = 99
		newMelts     = false
	)

	w, err := useCases.UpdateWidget(id, newName, newColor, newPrice, newInventory, newMelts)
	utils.CheckError(err)

	if w.Name != newName ||
		w.Color != newColor ||
		w.Price != newPrice ||
		w.Inventory != newInventory ||
		w.Melts != newMelts {
		t.Error("Should update widget with correct attributes")
	}
}
