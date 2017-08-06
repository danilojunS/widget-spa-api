package userRepo_test

import (
	user "github.com/danilojunS/widgets-spa-api/business/entities/user"
	userRepo "github.com/danilojunS/widgets-spa-api/business/repositories/user"
	utils "github.com/danilojunS/widgets-spa-api/utils"
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
	utils.CheckError(err)
}

func TestUserRead(t *testing.T) {
	users, err := userRepo.Read(createdUserID)
	utils.CheckError(err)

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
	utils.CheckError(err)

	if updatedUserID != createdUser.ID {
		t.Error("Should update correct user")
		return
	}

	users, err := userRepo.Read(createdUser.ID)
	utils.CheckError(err)

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
	utils.CheckError(err)

	if deletedUserID != updatedUser.ID {
		t.Error("Should delete correct user")
		return
	}

	users, err := userRepo.Read(deletedUserID)
	utils.CheckError(err)

	if len(users) > 0 {
		t.Error("Should not find deleted user")
	}
}

func TestClearUsers(t *testing.T) {
	users := []user.User{
		{Name: "user1", Gravatar: "gravatar1"},
		{Name: "user2", Gravatar: "gravatar2"},
		{Name: "user3", Gravatar: "gravatar3"},
	}

	for _, user := range users {
		_, err := userRepo.Create(user)
		utils.CheckError(err)
	}

	err := userRepo.Clear()
	utils.CheckError(err)

	usersFromRepo, err := userRepo.Read(0)
	utils.CheckError(err)

	if len(usersFromRepo) > 0 {
		t.Error("Should clear all users in repository")
	}
}

func TestTeardown(t *testing.T) {
	err := userRepo.Clear()
	utils.CheckError(err)
}
