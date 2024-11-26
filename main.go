package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now!")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("We have %v tickets remaining for this %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			names := strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of our bookings are: %v\n", firstNames)
	}

}
