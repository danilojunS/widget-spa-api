package useCase_test

import (
	widget "github.com/danilojunS/widgets-spa-api/business/entities/widget"
	widgetRepo "github.com/danilojunS/widgets-spa-api/business/repositories/widget"
	useCases "github.com/danilojunS/widgets-spa-api/business/use-cases"
	utils "github.com/danilojunS/widgets-spa-api/utils"
	"testing"
)

// Get Widgets without filtering
func TestGetWidgets(t *testing.T) {
	defer func() {
		err := widgetRepo.Clear()
		utils.CheckError(err)
	}()

	// create dummy widgets
	widgets := []widget.Widget{
		{Name: "widget1", Color: "color1", Price: "1.99"},
		{Name: "widget2", Color: "color2", Price: "2.99"},
		{Name: "widget3", Color: "color3", Price: "3.99"},
	}

	for _, widget := range widgets {
		_, err := widgetRepo.Create(widget)
		utils.CheckError(err)
	}

	widgetsFromUseCase, err := useCases.GetWidgets()
	utils.CheckError(err)

	if len(widgetsFromUseCase) != 3 {
		t.Error("Should get created widgets in use case")
	}
}
