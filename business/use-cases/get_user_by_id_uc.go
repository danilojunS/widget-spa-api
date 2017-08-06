package useCase

import (
	"errors"
	user "github.com/danilojunS/widgets-spa-api/business/entities/user"
	userRepo "github.com/danilojunS/widgets-spa-api/business/repositories/user"
)

// GetUserByID use case
func GetUserByID(id int) (user.User, error) {
	users, err := userRepo.Read(id)

	if err != nil {
		return user.User{}, err
	}

	if len(users) == 0 {
		return user.User{}, errors.New("No user found for ID")
	}

	return users[0], nil
}
