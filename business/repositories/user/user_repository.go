package user

import (
	"github.com/danilojunS/widgets-spa-api/business/entities/user"
)

var records []user.User
var lastInsertedID int

// Create a user
func Create(u user.User) (int, error) {
	u.ID = lastInsertedID
	lastInsertedID++
	records = append(records, u)
	return u.ID, nil
}

// Read users
func Read(id int) ([]user.User, error) {
	if id == 0 {
		return records, nil
	}

	var matches []user.User
	for _, r := range records {
		if r.ID == id {
			matches = append(matches, r)
		}
	}

	return matches, nil
}

// Update user
func Update(u user.User) (int, error) {
	for i, r := range records {
		if r.ID == u.ID {
			records[i] = u
		}
	}
	return u.ID, nil
}

// Delete user
func Delete(u user.User) (int, error) {
	for i, r := range records {
		if r.ID == u.ID {
			records = append(records[:i], records[i+1:]...)
		}
	}
	return u.ID, nil
}
