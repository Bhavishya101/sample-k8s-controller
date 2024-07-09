package main

import ("fmt")

// Entry point to the function
// Only one main function allowed

func main() {
	
	var confrenceName = "Go Confrence"
	const confrenceTickets = 50
	var remainingTickets = 50
	var bookings [50]string // Sample array 

	fmt.Printf("Welcome to %v booking application\n", confrenceName)
	fmt.Printf("We have a total of %v tickets and remaining Tickets are %v still available\n", confrenceTickets, remainingTickets )
	fmt.Printf("Get your tickets here to attend Confrence !!!!\n")
	//fmt.Println("Welcome to", confrenceName, "booking application !!!!")
	//fmt.Println("We have a total of", confrenceTickets, "tickets and", remainingTickets, "are still available")
	//fmt.Println("Get your tickets here to attend Confrence !!!!")

	//fmt.Println("Confrence Name is", confrenceName)
	fmt.Println(remainingTickets)
	fmt.Println(&remainingTickets) // Prints the Memory Address of the Pointer

	var firstName string
	var lastName string
	var userTickets int

	// userName = "Tom" 
	// userTickets = 2

	// Logic to allow user to ask for user input
	fmt.Println("Enter your name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("User %v %v has booked %v tickets. \n", firstName, lastName, userTickets)
	fmt.Printf("Number of Tickets remianing for confrence are %v. \n", remainingTickets)

	// arrays & slices
	// var bookings = [50]string{}  - Sample
	
	bookings[0] = firstName + " " + lastName
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n",bookings)
	fmt.Printf("Length of the array is %v\n", len(bookings))

}
