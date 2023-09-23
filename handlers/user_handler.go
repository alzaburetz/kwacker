package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/alzaburetz/kwacker/handlers/deserializers"
	"github.com/alzaburetz/kwacker/handlers/serializers"
	"github.com/alzaburetz/kwacker/models"
)

var users []models.User

func PostUser(w http.ResponseWriter, r *http.Request) {
	userString, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic("AAAAAAAAAAA")
	}

	user, err := deserializers.DeserializeUser(string(userString))

	if err == nil {
		users = append(users, user)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	usersString, err := serializers.SerializeUsers(users)

	if err != nil {
		panic("AAAAAAAAAAAAAA")
	}

	w.Write([]byte(usersString))
}
