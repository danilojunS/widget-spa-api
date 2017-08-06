package user_test

import (
	user "github.com/danilojunS/widgets-spa-api/business/entities/user"
	userRepo "github.com/danilojunS/widgets-spa-api/business/repositories/user"
	"testing"
)

// Fixtures for user data
const userFixtureName = "Danilo Jun"
const userFixtureGravatar = "https://d.js/gravatar"

var createdUserID int
var createdUser user.User
var updatedUser user.User

func TestUserCreate(t *testing.T) {
	u := user.User{Name: userFixtureName, Gravatar: userFixtureGravatar}
	var err error
	createdUserID, err = userRepo.Create(u)
	checkError(err)
}

func TestUserRead(t *testing.T) {
	users, err := userRepo.Read(createdUserID)
	checkError(err)

	if len(users) != 1 {
		t.Error("Should find one created user")
		return
	}

	createdUser = users[0]

	if createdUser.ID != createdUserID || createdUser.Name != userFixtureName || createdUser.Gravatar != userFixtureGravatar {
		t.Error("Should find created user")
	}
}

func TestUserUpdate(t *testing.T) {
	const newGravatar string = "https://d.js/awesome-gravatar"
	createdUser.Gravatar = newGravatar
	updatedUserID, err := userRepo.Update(createdUser)
	checkError(err)

	if updatedUserID != createdUser.ID {
		t.Error("Should update correct user")
		return
	}

	users, err := userRepo.Read(createdUser.ID)
	checkError(err)

	if len(users) != 1 {
		t.Error("Should find one updated user")
		return
	}

	updatedUser = users[0]

	if updatedUser.ID != createdUserID || updatedUser.Name != userFixtureName || updatedUser.Gravatar != newGravatar {
		t.Error("Should update user with correct attributes")
	}
}

func TestDeleteUser(t *testing.T) {
	deletedUserID, err := userRepo.Delete(updatedUser)
	checkError(err)

	if deletedUserID != updatedUser.ID {
		t.Error("Should delete correct user")
		return
	}

	users, err := userRepo.Read(deletedUserID)
	checkError(err)

	if len(users) > 0 {
		t.Error("Should not find deleted user")
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
