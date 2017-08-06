package useCase

import (
	user "github.com/danilojunS/widgets-spa-api/business/entities/user"
	userRepo "github.com/danilojunS/widgets-spa-api/business/repositories/user"
)

// GetUsers without filtering use case
func GetUsers() ([]user.User, error) {
	users, err := userRepo.Read(0)

	if err != nil {
		return nil, err
	}

	return users, nil
}
