package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"create_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}
	return Note{Title: title, Content: content, CreateAt: time.Now()}, nil
}

func (n Note) Display() {
	fmt.Printf("Your note titled %v has  the following content:\n\n%v\n\n", n.Title, n.Content)
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(n)
	if err != nil {
		return err
	}

	os.WriteFile(fileName, json, 0644)
	return nil
}
