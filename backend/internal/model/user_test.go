package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyID_WhenCreateNewUser_ShouldReceiveError(t *testing.T) {
	user := User{}
	assert.Error(t, user.IsValid(), "ID is required")
}

func TestGivenAnEmptyEmail_WhenCreateNewUser_ShouldReceiveError(t *testing.T) {
	user := NewUser("1", "", "1", "1")
	assert.Error(t, user.IsValid(), "username is required")
}

func TestGivenAnEmptyPassword_WhenCreateNewUser_ShouldReceiveError(t *testing.T) {
	user := NewUser("1", "a@b.c", "", "1")
	assert.Error(t, user.IsValid(), "password is required")
}

func TestGivenAValidUser_WhenCreateNewUser_ShouldNotReceiveError(t *testing.T) {
	user := NewUser("1", "a@b.c", "1", "1")
	assert.NoError(t, user.IsValid())
}
