package main

import (
	"booking-app/validator"
	"fmt"
)

type UserData struct {
	firstName      string
	lastName       string
	email          string
	countOfTickets uint
}

func main() {

	conferenceName := "GO conference" //declare and assign
	const totalTickets = 50
	var remainingTickets uint = 50
	var bookings []UserData

	greetingUser(conferenceName, totalTickets, remainingTickets)

	for {

		firstName, lastName, email, userTickets := getUserData()
		isValidUserData := validator.ValidateData(firstName, lastName, email)
		isValidTicketNumber := validator.ValidateTickets(remainingTickets, userTickets)

		if !isValidUserData || !isValidTicketNumber {
			continue
		}

		var userData = UserData{
			firstName:      firstName,
			lastName:       lastName,
			email:          email,
			countOfTickets: userTickets,
		}

		bookings = append(bookings, userData)

		fmt.Printf("Thank you %v %v. You booked %v tickets. your tickets send to your confirmed email %v\n", firstName, lastName, userTickets, email)

		remainingTickets = remainingTickets - userTickets
		fmt.Printf("%v tickets reamins", remainingTickets)

		var firstNames []string = getFirstNames(bookings)

		fmt.Printf("The list of bookings is: %v\n", firstNames)

		fmt.Printf("The bookings are: %v", bookings)

	}
}
