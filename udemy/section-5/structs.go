package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastname  string
	birthdate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthDate (MM/DD/YYYY): ")

	var appUser User

	appUser = User{
		firstName: userFirstName,
		lastname:  userLastName,
		birthdate: userBirthDate,
		createdAt: time.Now(),
	}

	outputUserDetails(&appUser)
}

func outputUserDetails(u *User) {
	fmt.Println(u.firstName, u.lastname, u.birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
