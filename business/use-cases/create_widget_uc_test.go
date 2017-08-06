package useCase_test

import (
	widgetRepo "github.com/danilojunS/widgets-spa-api/business/repositories/widget"
	useCases "github.com/danilojunS/widgets-spa-api/business/use-cases"
	utils "github.com/danilojunS/widgets-spa-api/utils"
	"testing"
)

// TestCreateWidget
func TestCreateWidget(t *testing.T) {
	defer func() {
		err := widgetRepo.Clear()
		utils.CheckError(err)
	}()

	const (
		name      = "My Widget"
		color     = "blue"
		price     = "9.99"
		inventory = 42
		melts     = true
	)

	w, err := useCases.CreateWidget(name, color, price, inventory, melts)
	utils.CheckError(err)

	if w.Name != name ||
		w.Color != color ||
		w.Price != price ||
		w.Inventory != inventory ||
		w.Melts != melts {
		t.Error("Should create widget with correct attributes")
	}
}
