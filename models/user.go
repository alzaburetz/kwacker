package models

import (
	"errors"
	"fmt"
	"strings"
)

const MIN_NAME_CHARACTERS int = 2

type User struct {
	Name            string
	Surname         string
	NickNameHandler string
	Posts           int
	IsValidated     bool
}

type UserModelValidator interface {
	ValidateName()
	ValidateSurname()
}

func (user User) ValidateName() error {
	if len(user.Name) < MIN_NAME_CHARACTERS {
		return fmt.Errorf("name must be at least %v elements", MIN_NAME_CHARACTERS)
	}

	firstLetter := string(user.Name[0])
	toUpperCase := strings.ToUpper(firstLetter)

	if firstLetter != toUpperCase {
		return errors.New("name must start with upper case")
	}

	return nil
}

func (user User) ValidateSurname() error {
	if len(user.Name) < MIN_NAME_CHARACTERS {
		return fmt.Errorf("surname must be at least %v elements", MIN_NAME_CHARACTERS)
	}

	firstLetter := string(user.Surname[0])
	toUpperCase := strings.ToUpper(firstLetter)

	if firstLetter != toUpperCase {
		return errors.New("surname must start with upper case")
	}

	return nil
}

func (user User) ValidateNick() error {
	if len(user.NickNameHandler) < MIN_NAME_CHARACTERS {
		return fmt.Errorf("nickname must be at least %v elements", MIN_NAME_CHARACTERS)
	}

	return nil
}

func NewUser(name string, surname string, nickname string, posts int) (*User, error) {
	if posts < 0 {
		return nil, errors.New("posts count cannot be less than zero")
	}

	var user User = User{
		Name:            name,
		Surname:         surname,
		NickNameHandler: nickname,
		Posts:           posts,
	}

	err := errors.Join(user.ValidateName(), user.ValidateSurname(), user.ValidateNick())

	if err != nil {
		return nil, err
	}

	user.IsValidated = true

	return &user, nil
}
