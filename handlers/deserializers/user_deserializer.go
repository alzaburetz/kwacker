package deserializers

import (
	"encoding/json"

	"github.com/alzaburetz/kwacker/models"
	"github.com/alzaburetz/kwacker/transport"
)

func DeserializeUser(userString string) (models.User, error) {
	var user transport.User

	err := json.Unmarshal([]byte(userString), &user)

	if err != nil {
		return models.User{}, err
	}

	domainUser, err := models.NewUser(user.Name, user.Surname, user.NickNameHandler, user.Posts)

	if err != nil {
		return models.User{}, err
	}

	return *domainUser, err
}
