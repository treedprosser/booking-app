package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50 // unsigned integer, can't be negative
// var bookings = [50]string{} // declare size, type and possibly some entries into the array
var bookings []string // alternate way to declare, no specific size makes it a slice (dynamic array)

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames() // saving function's return value as a variable
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is sold out. Come back next year!")
				break // terminates loop
			}
		} else {
			if !isValidName {
				fmt.Println("First or last name entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address must contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets entered is invalid")
			}
			//continue // skip this iteration and go to next one
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string { // input and output paramenters need a type
	firstNames := []string{}
	for _, booking := range bookings { // _ means an unused variable (index in this case)
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) { // functions can return multiple values
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their userName
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) // Scan is for input/output, & is a pointer, points to memory address of a variable

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	//bookings[0] = firstName + " " + lastName // assing a value to the array at the given index
	bookings = append(bookings, firstName+" "+lastName) // append value to the slice

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
