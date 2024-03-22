package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "GO conference" //declare and assign
	const totalTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("the type of conferenceName is %T, the type of the totaltickets is %T\n", conferenceName, totalTickets)

	fmt.Printf("Hello and welcome to %v. Here you can buy tickets\n", conferenceName)
	fmt.Printf("There is total %v tickets and %v are still available\n", totalTickets, remainingTickets)
	fmt.Println("Enjoy the conference!")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address")
		fmt.Scan(&email)

		fmt.Println("Enter the ticket numbers you want to book")
		fmt.Scan(&userTickets)

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v. You booked %v tickets. your tickets send to your confirmed email %v\n", firstName, lastName, userTickets, email)

		remainingTickets = remainingTickets - userTickets
		fmt.Printf("%v tickets reamins", remainingTickets)

		var firstNames []string
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The list of bookings is: %v\n", firstNames)
	}
}
