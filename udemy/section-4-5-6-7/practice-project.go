package main

import (
	"bufio"
	"fmt"
	"os"
	"projectfive.com/notes/note"
	"projectfive.com/notes/todo"
	"strings"
)

func getNoteData() (string, string) {
	title := getUserInput("note title: ")

	content := getUserInput("note content: ")

	return title, content
}

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")
	todoo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	outputData(todoo)

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(userNote)

	if err != nil {
		return
	}

}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("saving note failed")
		return err

	}
	fmt.Println("note successfully saved")
	return nil
}

func outputData(data outputtable) {
	data.Display()
	err := saveData(data)
	if err != nil {
		return
	}
}
