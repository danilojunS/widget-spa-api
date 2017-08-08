package user

import "errors"

// User business entity
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Gravatar string `json:"gravatar"`
}

// Validate a user
func (u User) Validate() error {
	if u.Name == "" {
		return errors.New("name is a required field and must not be empty string")
	}
	if u.Gravatar == "" {
		return errors.New("gravatar is a required field and must not be empty string")
	}

	return nil
}
