package userRepo

import (
	user "github.com/danilojunS/widgets-spa-api/business/entities/user"
	config "github.com/danilojunS/widgets-spa-api/config"
)

// Create a user
func Create(u user.User) (int, error) {
	if config.Get().DBMock {
		return CreateMock(u)
	}
	return CreatePg(u)
}

// Read users
func Read(id int) ([]user.User, error) {
	if config.Get().DBMock {
		return ReadMock(id)
	}
	return ReadPg(id)
}

// Update user
func Update(u user.User) (int, error) {
	if config.Get().DBMock {
		return UpdateMock(u)
	}
	return UpdatePg(u)
}

// Delete user
func Delete(u user.User) (int, error) {
	if config.Get().DBMock {
		return DeleteMock(u)
	}
	return DeletePg(u)
}

// Clear all users
func Clear() error {
	if config.Get().DBMock {
		return ClearMock()
	}
	return ClearPg()
}
