package main

import (
	"bufio"
	"fiverr.com/userInput/note"
	"fmt"
	"os"
	"strings"
)

func getNoteData() (string, string) {
	title := getUserInput("note title: ")

	content := getUserInput("note content: ")

	return title, content
}

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("saving note failed")
		return
	} else {
		fmt.Println("note successfully saved")
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
