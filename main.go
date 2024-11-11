package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	const remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now!")

	var userName string
	var userTickets int

	userName = "John Doe"
	userTickets = 2

	fmt.Printf("User %v has booked %v tickets\n", userName, userTickets)

}
