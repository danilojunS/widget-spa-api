package scripts

import (
	user "github.com/danilojunS/widgets-spa-api/business/entities/user"
	userRepo "github.com/danilojunS/widgets-spa-api/business/repositories/user"
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
