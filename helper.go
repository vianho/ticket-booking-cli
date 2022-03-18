package main

import "strings"

// global variable
var Global = "global variable"

// export a function by capitalizing the first letter
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint8, remainingTickets uint8) (bool, bool, bool) {
	// validate name
	isValidName := len(firstName) > 2 && len(lastName) > 2
	// validate email
	isValidEmail := strings.Contains(email, "@")
	// validate number of tickets
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTickets
}
