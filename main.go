// main file

package main // all code must belong to a package

// we need to tell go where is the entry point of the program

import "fmt"

// import fmt package

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application!")
	fmt.Println("We have total of", conferenceTickets, "tickets and",
		remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend!")
}
