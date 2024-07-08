package main

import ("fmt")

// Entry point to the function
// Only one main function allowed

func main() {
	
	var confrenceName = "Go Confrence"
	const confrenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", confrenceName)
	fmt.Printf("We have a total of %v tickets and remaining Tickets are %v still available\n", confrenceTickets, remainingTickets )
	fmt.Printf("Get your tickets here to attend Confrence !!!!\n")
	//fmt.Println("Welcome to", confrenceName, "booking application !!!!")
	//fmt.Println("We have a total of", confrenceTickets, "tickets and", remainingTickets, "are still available")
	//fmt.Println("Get your tickets here to attend Confrence !!!!")

	//fmt.Println("Confrence Name is", confrenceName)

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2
	//remainingTickets = remainingTickets - userTickets

	fmt.Printf("User %v booked %v tickets. \n", userName, userTickets)
}
