package useCase_test

import (
	user "github.com/danilojunS/widgets-spa-api/business/entities/user"
	userRepo "github.com/danilojunS/widgets-spa-api/business/repositories/user"
	useCases "github.com/danilojunS/widgets-spa-api/business/use-cases"
	utils "github.com/danilojunS/widgets-spa-api/utils"
	"testing"
)

// Get Users without filtering
func TestGetUsers(t *testing.T) {
	defer func() {
		err := userRepo.Clear()
		utils.CheckError(err)
	}()

	// create dummy users
	users := []user.User{
		{Name: "Danilo", Gravatar: "https://danilo.com/gravatar"},
		{Name: "Jun", Gravatar: "https://jun.com/gravatar"},
		{Name: "Shibata", Gravatar: "https://shibata.com/gravatar"},
	}

	for _, user := range users {
		_, err := userRepo.Create(user)
		utils.CheckError(err)
	}

	usersFromUseCase, err := useCases.GetUsers()
	utils.CheckError(err)

	if len(usersFromUseCase) != 3 {
		t.Error("Should get created users in use case")
	}
}
