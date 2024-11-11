// main file

package main // all code must belong to a package

// we need to tell go where is the entry point of the program

import "fmt"

// import fmt package

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}
