package user_test

import (
	"github.com/danilojunS/widgets-spa-api/business/entities/user"
	"testing"
)

// Fixtures for user data
var userFixtureID = 1
var userFixtureName = "Danilo Jun"
var userFixtureGravatar = "https://d.js/gravatar"

func TestCreateUser(t *testing.T) {
	u := user.User{ID: userFixtureID, Name: userFixtureName, Gravatar: userFixtureGravatar}
	if err := u.Validate(); err != nil {
		t.Error("Should create a valid user with ID, Name and Gravatar")
	}
}

func TestCreateUserFailMissingId(t *testing.T) {
	u := user.User{Name: userFixtureName, Gravatar: userFixtureGravatar}
	if err := u.Validate(); err == nil {
		t.Error("Should not create a valid user without ID")
	}
}

func TestCreateUserFailMissingName(t *testing.T) {
	u := user.User{ID: 1, Gravatar: userFixtureGravatar}
	if err := u.Validate(); err == nil {
		t.Error("Should not create a valid user without Name")
	}
}

func TestCreateUserFailMissingGravatar(t *testing.T) {
	u := user.User{ID: 1, Name: userFixtureName}
	if err := u.Validate(); err == nil {
		t.Error("Should not create a valid user without Gravatar")
	}
}
