package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("What's Your Name?")
	println("Fuck you", userFirstName)

	userLastName := getUserData("Now your last name?")
	userBirthdate := getUserData("please enter your birthdate (MM/DD/YYYY): ")

	var appUser User
	appUser = User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	outputUserDetails(firstName, lastName, birthdate)
}

func outputUserDetails(firstName, lastName, birthdate string) {
	// ..
	fmt.Println(firstName, lastName, birthdate)

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
