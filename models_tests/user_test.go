package models_test

import (
	"testing"

	user "github.com/alzaburetz/kwacker/models"
)

// ValidateName method should work correctly if the name is correct
func TestValidateName(t *testing.T) {
	// Arrange
	var user user.User = user.User{Name: "John"}

	// Act
	err := user.ValidateName()

	// Assert
	if err != nil {
		t.Errorf("Validation didn't work properly")
	}
}

// ValidateName method should return error if name is lower case
func TestValidateNameShouldReturnError(t *testing.T) {
	// Arrange
	var user user.User = user.User{Name: "john"}

	// Act
	err := user.ValidateName()

	// Assert
	if err == nil {
		t.Errorf("Validation complete wrongfully")
	}
}

// New user can be created
func TestNewUserCanBeCreated(t *testing.T) {
	// Act
	user, err := user.NewUser("John", "Dow", "john_doe", 2)

	// Assert
	if user == nil || err != nil {
		t.Errorf("Could not create a new user")
	}
}
