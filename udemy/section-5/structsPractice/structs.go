package main

import (
	"5.com/structs-practice/structsPractice"
	"fmt"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthDate (MM/DD/YYYY): ")

	var appUser *structsPractice.User
	appUser, err := structsPractice.NewUser(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}
	admin := structsPractice.NewAdmin("test@example.com", "test123")
	admin.OutputUserDetails()
	admin.ClearUserData()
	appUser.OutputUserDetails()
	appUser.ClearUserData()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
