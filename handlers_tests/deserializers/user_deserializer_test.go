package deserializers_test

import (
	"testing"

	"github.com/alzaburetz/kwacker/handlers/deserializers"
	"github.com/alzaburetz/kwacker/models"
)

// Deserializer works correctly
func TestUserDeserializerWorksCorrectly(t *testing.T) {
	// Arrange
	jsonString := "{\"name\":\"John\",\"surname\":\"Doe\",\"posts\":2,\"nick_name\":\"john_doe\"}"
	expectedUser := models.User{
		Name:            "John",
		Surname:         "Doe",
		Posts:           2,
		NickNameHandler: "john_doe",
		IsValidated:     true,
	}

	// Act
	user, err := deserializers.DeserializeUser(jsonString)

	// Assert
	if user != expectedUser || err != nil {
		t.Errorf("Error while decodnig json string to a model")
	}
}

func TestUserDeserializerFailsToDeserializeInvalidData(t *testing.T) {
	// Arrange
	jsonString := "{\"name\":\"John\",\"surname\":\"Doe\",\"posts\":-6,\"nick_name\":\"john_doe\"}"

	// Act
	_, err := deserializers.DeserializeUser(jsonString)

	// Assert
	if err != nil {
		return
	}

	t.Errorf("Error while decodnig json string to a model")
}
