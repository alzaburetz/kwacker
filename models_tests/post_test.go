package models_test

import (
	"testing"
	"time"

	"github.com/alzaburetz/kwacker/models"
)

func TestNewBlockWorksCorrectly(t *testing.T) {
	// Act
	block, err := models.NewBlock(models.Text, "content", "style")

	// Assert
	if err != nil {
		t.Errorf("Could not created block: %v", block)
	}
}

func TestNewBlockReturnsNilIfContentStringIsEmpty(t *testing.T) {
	// Act
	_, err := models.NewBlock(models.Text, "", "style")

	// Assert
	if err != nil {
		return
	}

	t.Error("Created block with invalid data")
}

func TestNewBlockReturnsNilIfStyleStringIsEmpty(t *testing.T) {
	// Act
	_, err := models.NewBlock(models.Text, "content", "")

	// Assert
	if err != nil {
		return
	}

	t.Error("Created block with invalid data")
}

func TestNewPostWorksCorrectly(t *testing.T) {
	// Arrange
	block, _ := models.NewBlock(models.Text, "content", "style")
	var blocks []models.PostBlock
	blocks = append(blocks, *block)

	// Act
	post, err := models.NewPost("headline", time.Now(), time.Now(), blocks)

	// Arrange
	if err != nil {
		t.Errorf("Could not create a post: %v", post)
	}
}

func TestNewPostsReturnsNilIfHeadlineIsEmpty(t *testing.T) {
	// Arrange
	block, _ := models.NewBlock(models.Text, "content", "style")
	var blocks []models.PostBlock
	blocks = append(blocks, *block)

	// Act
	post, err := models.NewPost("", time.Now(), time.Now(), blocks)

	// Arrange
	if err != nil {
		return
	}

	t.Errorf("Still created a post: %v", post)
}

func TestNewPostsReturnsNilIfCreatedDateIsInvalid(t *testing.T) {
	// Arrange
	block, _ := models.NewBlock(models.Text, "content", "style")
	var blocks []models.PostBlock
	blocks = append(blocks, *block)

	// Act
	post, err := models.NewPost("headline", time.Now().Add(time.Hour), time.Now(), blocks)

	// Arrange
	if err != nil {
		return
	}

	t.Errorf("Still created a post: %v", post)
}

func TestNewPostsReturnsNilIfEditedDateIsInvalid(t *testing.T) {
	// Arrange
	block, _ := models.NewBlock(models.Text, "content", "style")
	var blocks []models.PostBlock
	blocks = append(blocks, *block)

	// Act
	post, err := models.NewPost("headline", time.Now(), time.Now().Add(time.Hour), blocks)

	// Arrange
	if err != nil {
		return
	}

	t.Errorf("Still created a post: %v", post)
}

func TestNewPostsReturnsNilIfAnyBlockIsNotValidated(t *testing.T) {
	// Arrange
	var blocks []models.PostBlock
	blocks = append(blocks, models.PostBlock{})

	// Act
	post, err := models.NewPost("headline", time.Now(), time.Now(), blocks)

	// Arrange
	if err != nil {
		return
	}

	t.Errorf("Still created a post: %v", post)
}
