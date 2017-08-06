package useCase_test

import (
	user "github.com/danilojunS/widgets-spa-api/business/entities/user"
	userRepo "github.com/danilojunS/widgets-spa-api/business/repositories/user"
	useCases "github.com/danilojunS/widgets-spa-api/business/use-cases"
	utils "github.com/danilojunS/widgets-spa-api/utils"
	"testing"
)

// Get Users by ID
func TestGetUserByID(t *testing.T) {
	defer func() {
		err := userRepo.Clear()
		utils.CheckError(err)
	}()

	const userName = "Danilo"
	const userGravatar = "https://danilo.com/gravatar"

	// create dummy user
	user := user.User{Name: userName, Gravatar: userGravatar}
	id, err := userRepo.Create(user)
	utils.CheckError(err)

	userFromUseCase, err := useCases.GetUserByID(id)
	utils.CheckError(err)

	if userFromUseCase.Name != userName || userFromUseCase.Gravatar != userGravatar {
		t.Error("Should get created user by ID in use case")
	}
}
