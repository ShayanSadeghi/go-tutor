package validator

import (
	"fmt"
	"strings"
)

// exporting function from this package by capitalize function names :)
func ValidateData(firstName string, lastName string, email string) bool {
	isFirstNameValid := len(firstName) >= 2
	isLastNameValid := len(lastName) >= 2
	isEmailValid := strings.Contains(email, "@")

	if !isFirstNameValid {
		fmt.Println("First name should be at least two characters!")
	}

	if !isLastNameValid {
		fmt.Println("Last name should be at least two characters!")
	}

	if !isEmailValid {
		fmt.Println("Email address should contain '@' character!")
	}

	return isFirstNameValid && isLastNameValid && isEmailValid

}

func ValidateTickets(remainingTickets uint, userTickets uint) bool {
	isValidCount := remainingTickets > userTickets

	if !isValidCount {
		fmt.Printf("There's only %v tickets remaining, so you cannot book %v tickets!\n", remainingTickets, userTickets)
	}

	return isValidCount

}
