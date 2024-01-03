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
	Title     string    `json:"title"` // json tag -- json package will use this tag to encode and decode json
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("your note title %v has the following content: \n\n%v\n\n", note.Title, note.Content)

}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return os.WriteFile(fileName, json, 0644)

}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {

		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
