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

// New user can not be created if name starts with lower-case letter
func TestNewUserCanNotBeCreatedIfNameStartsWithLowerCase(t *testing.T) {
	// Act
	user, err := user.NewUser("john", "Doe", "john_doe", 2)

	// Assert
	if user != nil || err.Error() != "name must start with upper case" {
		t.Errorf("User still got created with invalid name")
	}
}

// New user can not be created if surname starts with lower-case letter
func TestNewUserCanNotBeCreatedIfSurnameStartsWithLowerCase(t *testing.T) {
	// Act
	user, err := user.NewUser("John", "doe", "john_doe", 2)

	// Assert
	if user != nil || err.Error() != "surname must start with upper case" {
		t.Errorf("User still got created with invalid surname")
	}
}

// New user can not be created if nickname consist of less than 2 characters
func TestNewUserCanNotBeCreatedIfNickNameIsInvalid(t *testing.T) {
	// Act
	user, err := user.NewUser("John", "Doe", "a", 2)

	// Assert
	if user != nil || err.Error() != "nickname must be at least 2 elements" {
		t.Errorf("User still got created with invalid nickname")
	}
}

func TestNewUserCanNotBeCreatedIfPostsAmountIsLessThanZero(t *testing.T) {
	// Act
	user, err := user.NewUser("John", "Doe", "john_doe", -1)

	// Assert
	if user != nil || err.Error() != "posts count cannot be less than zero" {
		t.Errorf("User still got created with invalid posts amount")
	}
}
