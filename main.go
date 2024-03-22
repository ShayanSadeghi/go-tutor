package main

import (
	"booking-app/validator"
	"fmt"
	"sync"
	"time"
)

type UserData struct {
	firstName      string
	lastName       string
	email          string
	countOfTickets uint
}

var wg = sync.WaitGroup{} // we use this to prevent code exiting before all threads done

func main() {

	conferenceName := "GO conference" //declare and assign
	const totalTickets = 50
	var remainingTickets uint = 50
	var bookings []UserData

	greetingUser(conferenceName, totalTickets, remainingTickets)

	// for {

	firstName, lastName, email, userTickets := getUserData()
	isValidUserData := validator.ValidateData(firstName, lastName, email)
	isValidTicketNumber := validator.ValidateTickets(remainingTickets, userTickets)

	if !isValidUserData || !isValidTicketNumber {
		fmt.Println("There is an error in user data")
		// continue
	}

	var userData = UserData{
		firstName:      firstName,
		lastName:       lastName,
		email:          email,
		countOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	wg.Add(1) // as we create a new thread in the following line, we add 1 to the wg counter
	go sendTicket(userTickets, firstName, lastName, email)

	fmt.Printf("Thank you %v %v. You booked %v tickets. your tickets send to your confirmed email %v\n", firstName, lastName, userTickets, email)

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("%v tickets reamins", remainingTickets)

	var firstNames []string = getFirstNames(bookings)

	fmt.Printf("The list of bookings is: %v\n", firstNames)

	fmt.Printf("The bookings are: %v\n", bookings)

	wg.Wait() // wait until the counter become 0
	// }
}

func sendTicket(ticketsNumber uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second) // this process takes time to finish

	var ticketData = fmt.Sprintf("%v tickets booked for %v %v\n", ticketsNumber, firstName, lastName) // save formatted text in strting variable

	fmt.Println("###########")
	fmt.Printf("Ticket\n%v\nis sending to %v\n", ticketData, email)
	fmt.Println("###########")
	wg.Done() // decrease 1 from the wg counter
}
