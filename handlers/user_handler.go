package handlers

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/alzaburetz/kwacker/handlers/deserializers"
	"github.com/alzaburetz/kwacker/models"
	"github.com/alzaburetz/kwacker/storage"
)

var users []models.User

func AddUser(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	fmt.Printf("Recieved: %v\n", string(body))

	user, _ := deserializers.DeserializeUser(string(body))

	fmt.Printf("Deserialized as: %v\n", user)

	userModel, _ := models.NewUser(user.Name, user.Surname, user.NickNameHandler, 0)

	ctx := context.Background()

	defer ctx.Done()

	_, err := storage.AddUser(&ctx, *userModel)

	if err != nil {
		fmt.Printf("ERROR OCCURED: %v\n", err.Error())
		return
	}

	fmt.Println("User added")
}
