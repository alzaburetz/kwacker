package serializers

import (
	"encoding/json"

	"github.com/alzaburetz/kwacker/models"
	"github.com/alzaburetz/kwacker/transport"
)

func SerializeUser(user models.User) (string, error) {
	if !user.IsValidated {
		panic("User must be validated")
	}

	transportUser := transport.User{
		Name:            user.Name,
		Surname:         user.Surname,
		Posts:           user.Posts,
		NickNameHandler: user.NickNameHandler,
		Email:           user.Email,
	}
	result, err := json.Marshal(transportUser)

	return string(result), err
}

func SerializeUsers(users []models.User) (string, error) {
	for _, user := range users {
		if !user.IsValidated {
			panic("User must be validated")
		}
	}

	result, err := json.Marshal(users)

	return string(result), err
}
