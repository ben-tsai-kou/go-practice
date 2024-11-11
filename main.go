// main file

package main // all code must belong to a package

// we need to tell go where is the entry point of the program

import "fmt"

// import fmt package

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

	var userName string = ""
	var userTicket int = 0
	// ask user for their name

	userName = "Tom"
	userTicket = 2

	fmt.Printf("User %v booked %v tickets.\n", userName, userTicket)
}
