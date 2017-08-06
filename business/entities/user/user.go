package user

import "errors"

// User business entity
type User struct {
	ID       int
	Name     string
	Gravatar string
}

// Validate a user
func (u User) Validate() error {
	if u.ID == 0 {
		return errors.New("id is a required field and must not be 0")
	}
	if u.Name == "" {
		return errors.New("name is a required field and must not be empty string")
	}

	if u.Gravatar == "" {
		return errors.New("gravatar is a required field and must not be empty string")
	}

	return nil
}
