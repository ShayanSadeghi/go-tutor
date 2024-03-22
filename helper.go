package main

import (
	"fmt"
	"strings"
)

func greetingUser(confName string, totalTickets uint, remainingTickets uint) {
	fmt.Printf("the type of conferenceName is %T, the type of the totaltickets is %T\n", confName, totalTickets)

	fmt.Printf("Hello and welcome to %v. Here you can buy tickets\n", confName)
	fmt.Printf("There is total %v tickets and %v are still available\n", totalTickets, remainingTickets)
	fmt.Println("Enjoy the conference!")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserData() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}
