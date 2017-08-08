package scripts

import (
	user "github.com/danilojunS/widgets-spa-api/business/entities/user"
	widget "github.com/danilojunS/widgets-spa-api/business/entities/widget"
	userRepo "github.com/danilojunS/widgets-spa-api/business/repositories/user"
	widgetRepo "github.com/danilojunS/widgets-spa-api/business/repositories/widget"
	utils "github.com/danilojunS/widgets-spa-api/utils"
)

// UserPopulate populates dummy users
func UserPopulate() {
	users := []user.User{
		{Name: "Danilo", Gravatar: "https://danilo/gravatar"},
		{Name: "Jun", Gravatar: "https://jun/gravatar"},
		{Name: "Shibata", Gravatar: "https://shibata/gravatar"},
	}

	for _, user := range users {
		_, err := userRepo.Create(user)
		utils.CheckError(err)
	}
}

// WidgetPopulate populates dummy widgets
func WidgetPopulate() {
	widgets := []widget.Widget{
		{
			Name:      "My widget",
			Color:     "blue",
			Price:     "9.99",
			Inventory: 42,
			Melts:     true,
		},
		{
			Name:      "My other widget",
			Color:     "red",
			Price:     "19.99",
			Inventory: 99,
			Melts:     false,
		},
	}

	for _, widget := range widgets {
		_, err := widgetRepo.Create(widget)
		utils.CheckError(err)
	}
}
