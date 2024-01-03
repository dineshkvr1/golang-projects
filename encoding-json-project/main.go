package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"structs/encoding-json-project/note"
)

func main() {

	title, content := getNoteData()

	//new constructor function
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

	err = userNote.Save()

	if err != nil {
		fmt.Println("saving the note failed ")
		return
	}

	fmt.Println("saving the note was successful")

}

func getNoteData() (string, string) {
	title := getUserInput("note title: ")

	content := getUserInput("note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v", prompt)
	// getting longer text input (read)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		//fmt.Println(err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

// it will create json file with title name, content and created_at
// go run main.go
// note title: my first note
// note content: this is my first note
// your note title my first note has the following content:
