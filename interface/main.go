package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"structs/interface/note"
	"structs/interface/todo"
)

type saver interface {
	Save() error
}

//type displayer interface {
//	Display()
//}

type outputtable interface {
	saver
	Display()
}

func main() {

	printSomething(1)
	printSomething("hello")

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}
	//new constructor function
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	// interface usage and implementation
	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)

}

// any kind of value accept using this special interface {}
func printSomething(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("this is an int", value)
	case string:
		fmt.Println("this is a string", value)
	case float64:
		fmt.Println("this is a float64", value)
	default:
		fmt.Println("this is something else", value)
	}
	//fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

// interface usage and implementation
func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("saving the note failed ")
		return err
	}

	fmt.Println("saving the note was successful")
	return nil

}

func getTodoData() string {
	return getUserInput("Todo text: ")
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
