package books

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

const (
	CategoryNovel      = "CategoryNovel"
	CategoryShortStory = "CategoryShortStory"
)

var (
	ErrInvalidJSON    = errors.New("JSON is invalid")
	ErrIncompleteJSON = errors.New("JSON is incomplete")
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func (b *Book) AuthorFirstName() string {
	// Assuming author's name is composed by: "firstName lastName"
	nameAndLastName := strings.Split(b.Author, " ")
	if len(nameAndLastName) < 2 {
		return ""
	}

	return nameAndLastName[0]
}

func (b *Book) AuthorLastName() string {
	// Assuming author's name is composed by: "firstName lastName"
	nameAndLastName := strings.Split(b.Author, " ")
	if len(nameAndLastName) < 2 {
		return nameAndLastName[0]
	}

	return nameAndLastName[len(nameAndLastName)-1]
}

func (b *Book) IsValid() bool {
	return b.Author != ""
}

func (b *Book) Category() string {
	if b.Pages > 300 {
		return CategoryNovel
	} else {
		return CategoryShortStory
	}
}

func (b *Book) AsJSON() (string, error) {
	jsonData, err := json.Marshal(b)
	if err != nil {
		return "", fmt.Errorf("error encoding to JSON: %w", err)
	}

	return string(jsonData), nil
}

func NewBookFromJSON(data string) (*Book, error) {
	var book Book
	err := json.Unmarshal([]byte(data), &book)
	if err != nil {
		var syntaxErr *json.SyntaxError
		if errors.As(err, &syntaxErr) {
			return nil, ErrInvalidJSON
		}

		return nil, fmt.Errorf("error decoding from JSON: %w", err)
	}

	// Check for required fields
	if book.Title == "" || book.Author == "" || book.Pages == 0 {
		return nil, ErrIncompleteJSON
	}

	return &book, nil
}
