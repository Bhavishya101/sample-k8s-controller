package main

import (
	"fmt"
	"strings"
)

// Entry point to the function
// Only one main function allowed

const confrenceTickets = 50
var confrenceName = "Go Confrence"
var remainingTickets = 50
//var bookings [50]string // Sample array 
var bookings = []string{}


func main() {
	
	// Since we are using function now
	// fmt.Printf("Welcome to %v booking application\n", confrenceName)
	// fmt.Printf("We have a total of %v tickets and remaining Tickets are %v still available\n", confrenceTickets, remainingTickets )
	// fmt.Printf("Get your tickets here to attend Confrence !!!!\n")


	//fmt.Println("Welcome to", confrenceName, "booking application !!!!")
	//fmt.Println("We have a total of", confrenceTickets, "tickets and", remainingTickets, "are still available")
	//fmt.Println("Get your tickets here to attend Confrence !!!!")

	//fmt.Println("Confrence Name is", confrenceName)
	// fmt.Println(remainingTickets)
	// fmt.Println(&remainingTickets) // Prints the Memory Address of the Pointer

	for {
		greetUsers()

		// for emai we can write below
		// isvalidemail := strings.Contains(email, "@")
		// Sample for checking city for booking ticket
		// isValidCity := city == "Singapore" || city == "London"
		firstName, lastName, userTickets := getUserInput()

		isValidName, isValidTicketNumber :=  validateUserInput(firstName, lastName, userTickets, remainingTickets)

		if isValidName && isValidTicketNumber {

			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName)
	
			firstNames := getFirstNames()
			fmt.Printf("These first names of the bookings: %v\n", firstNames)
	
			if remainingTickets == 0 {
				// end the programme
				fmt.Println("Our confrence is fully booked, come back next year")
				break
			}
		} else if userTickets == remainingTickets {
			fmt.Println("You are booking all our tickets, thankyou")
			break
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidTicketNumber{
				fmt.Println("Not valid ticket number")
			}
			fmt.Println("You have entered an incorrect input")
			continue
		}

	}

	// arrays & slices
	// var bookings = [50]string{}  - Sample
	
	// bookings[0] = firstName + " " + lastName

	// fmt.Printf("The first value: %v\n", bookings[0]
	// fmt.Printf("Array type: %T\n",bookings)
	// fmt.Printf("Length of the array is %v\n", len(bookings))
}


func greetUsers() {
	fmt.Printf("Welcome to %v , we have total %v, out of which remining tickets are %v!!!\n", confrenceName , confrenceTickets, remainingTickets)
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings { // _ is a blank identifier,so that we can define a varibale an not use it
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, userTickets int, remainingTickets int ) (bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >=2 
	isValidTicketNumber := userTickets> 0 && userTickets <= remainingTickets
	return isValidName, isValidTicketNumber
}

func getUserInput() (string, string, int) {
	var firstName string
	var lastName string
	var userTickets int

	fmt.Println("Enter your name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, userTickets
}

func bookTicket(remainingTickets int, userTickets int, bookings []string, firstName string, lastName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("User %v %v has booked %v tickets. \n", firstName, lastName, userTickets)
	fmt.Printf("Number of Tickets remianing for confrence are %v. \n", remainingTickets)
}