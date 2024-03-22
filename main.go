package main

import (
	"booking-app/validator"
	"fmt"
	"strconv"
)

func main() {
	conferenceName := "GO conference" //declare and assign
	const totalTickets = 50
	var remainingTickets uint = 50
	var bookings = make([]map[string]string, 0) // a slice(list) of maps, with initially 0 length

	greetingUser(conferenceName, totalTickets, remainingTickets)

	for {

		firstName, lastName, email, userTickets := getUserData()
		isValidUserData := validator.ValidateData(firstName, lastName, email)
		isValidTicketNumber := validator.ValidateTickets(remainingTickets, userTickets)

		if !isValidUserData || !isValidTicketNumber {
			continue
		}

		var userData = make(map[string]string) // [key type]value type

		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["countOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

		bookings = append(bookings, userData)

		fmt.Printf("Thank you %v %v. You booked %v tickets. your tickets send to your confirmed email %v\n", firstName, lastName, userTickets, email)

		remainingTickets = remainingTickets - userTickets
		fmt.Printf("%v tickets reamins", remainingTickets)

		var firstNames []string = getFirstNames(bookings)

		fmt.Printf("The list of bookings is: %v\n", firstNames)
	}
}
