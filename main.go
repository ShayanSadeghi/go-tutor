package main

import (
	"fmt"
)

func main() {
	conferenceName := "GO conference" //declare and assign
	const totalTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	greetingUser(conferenceName, totalTickets, remainingTickets)

	for {

		firstName, lastName, email, userTickets := getUserData()

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v. You booked %v tickets. your tickets send to your confirmed email %v\n", firstName, lastName, userTickets, email)

		remainingTickets = remainingTickets - userTickets
		fmt.Printf("%v tickets reamins", remainingTickets)

		var firstNames []string = getFirstNames(bookings)

		fmt.Printf("The list of bookings is: %v\n", firstNames)
	}
}
