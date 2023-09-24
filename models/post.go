package models

import (
	"errors"
	"time"
)

type Post struct {
	HeadLine    string
	Created     time.Time
	Edited      time.Time
	Blocks      []PostBlock
	IsValidated bool
}

type PostBlock struct {
	Type        PostBlockType
	Content     string
	Style       string
	IsValidated bool
}

type PostBlockType int64

const (
	None PostBlockType = iota
	Text
	Image
)

func ValidateString(str string) error {
	if len(str) == 0 {
		return errors.New("string should not be empty")
	}

	return nil
}

func ValidateTime(datetime time.Time) error {
	if datetime.Unix() > time.Now().Unix() {
		return errors.New("date cannot be in future")
	}

	return nil
}

func NewPost(headline string, created time.Time, edited time.Time, blocks []PostBlock) (*Post, error) {
	for _, block := range blocks {
		if !block.IsValidated {
			return nil, errors.New("one ore more blocks is not validated")
		}
	}

	err := errors.Join(
		ValidateString(headline),
		ValidateTime(created),
		ValidateTime(edited),
	)

	if err != nil {
		return nil, err
	}

	post := Post{
		HeadLine:    headline,
		Created:     created,
		Edited:      edited,
		Blocks:      blocks,
		IsValidated: true,
	}

	return &post, nil
}

func NewBlock(blockType PostBlockType, content string, style string) (*PostBlock, error) {
	err := errors.Join(
		ValidateString(content),
		ValidateString(style),
	)

	if err != nil {
		return nil, err
	}

	block := PostBlock{
		Type:        blockType,
		Content:     content,
		Style:       style,
		IsValidated: true,
	}

	return &block, nil
}
