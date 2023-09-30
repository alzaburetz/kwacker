package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/alzaburetz/kwacker/models"
	storage_models "github.com/alzaburetz/kwacker/storage/models"
	"github.com/go-pg/pg/v10/orm"
)

func AddUser(ctx *context.Context, user models.User) (*orm.Result, error) {
	connection := GetDatabase().Conn()

	fmt.Printf("Connection established: %v\n", connection == nil)

	defer connection.Close()

	storageUser := storage_models.User{
		Name:           user.Name,
		Surname:        user.Surname,
		Created:        time.Now(),
		Edited:         time.Now(),
		ProfilePicture: "",
		Email:          "",
	}

	fmt.Printf("Storage model created: %v\n", storageUser)

	result, err := connection.Model(&storageUser).Insert()

	fmt.Printf("%v\n", result)

	if err != nil {
		panic(err)
	}

	return &result, nil
}
