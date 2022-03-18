package main

import (
	"fmt"
	"sync"
	"time"
)

// package level variables
const exhibitionName string = "Attack On Titan: The Exhibition"
const exhibitionTickets uint8 = 50

var remainingTickets uint8 = exhibitionTickets
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint8
}

// wait for all goroutines to finish
// use wait group in 3 spots:
// 1. main()
//    - wg.Add(1) to add goroutine to wait group
// 		- wg.Wait() to wait for all goroutines to finish
// 2. bookTickets()
//    - wg.Done() to indicate goroutine is finished
var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	// for remainingTickets > 0 && len(bookings) < 50 {

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTickets := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTickets {

		bookTickets(userTickets, firstName, lastName, email)
		// add go keyword for concurrency
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("Bookings: %v\n", firstNames)

		fmt.Println("==========================================================================")

		if remainingTickets == 0 {
			fmt.Println("Sorry, we are sold out!")
			fmt.Println("==========================================================================")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("First name or last name too short!")
		}
		if !isValidEmail {
			fmt.Println("Invalid email!")
		}
		if !isValidTickets {
			fmt.Println("Invalid number of tickets!")
		}
		fmt.Println("==========================================================================")
	}
	wg.Wait()

	// }
}

func greetUsers() {
	fmt.Printf("Welcome to %s booking application!\n", exhibitionName)
	fmt.Printf("There are %d/%d tickets available.\n", remainingTickets, exhibitionTickets)
	fmt.Println("Get your tickets here to attend!")
}

func getUserInput() (string, string, string, uint8) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint8

	// ask user for name
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	// ask user for email
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	// ask user for number of tickets
	fmt.Print("How many tickets would you like to buy? ")
	fmt.Scanf("%d", &userTickets)
	fmt.Printf("You have requested %d ticket(s).\n", userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint8, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %s for booking %d tickets! You will receive a confirmation email at %s\n", firstName, userTickets, email)
	fmt.Printf("There are %d/%d tickets available for %s.\n", remainingTickets, exhibitionTickets, exhibitionName)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func sendTicket(userTickets uint8, firstName string, lastName string, email string) {
	time.Sleep(20 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##########################################################################")
	fmt.Printf("Sending ticket:\n%v \nto email address: %v\n", ticket, email)
	fmt.Println("##########################################################################")
	wg.Done()
}
