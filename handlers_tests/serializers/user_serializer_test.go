package serializers_test

import (
	"testing"

	"github.com/alzaburetz/kwacker/handlers/serializers"
	"github.com/alzaburetz/kwacker/models"
)

func TestSerializeUserWorksCorrectly(t *testing.T) {
	// Arrange
	user, _ := models.NewUser("John", "Doe", "john", "test@test.com", 2)
	expectedJson := "{\"name\":\"John\",\"surname\":\"Doe\",\"nick_name\":\"john\",\"posts\":2,\"email\":\"test@test.com\"}"

	// Act
	result, err := serializers.SerializeUser(*user)

	// Assert
	if result != expectedJson || err != nil {
		t.Errorf("Wanted: %v, Got: %v", expectedJson, result)
	}
}
