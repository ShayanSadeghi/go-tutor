package main

import "fmt"

func main() {
	const conferenceName = "GO conference"
	const totalTickets = 50
	var remainingTickets = 50

	fmt.Printf("the type of conferenceName is %T, the type of the totaltickets is %T\n", conferenceName, totalTickets)

	fmt.Printf("Hello and welcome to %v. Here you can buy tickets\n", conferenceName)
	fmt.Printf("There is total %v tickets and %v are still available\n", totalTickets, remainingTickets)
	fmt.Println("Enjoy the conference!")

	var userName string
	var userTickets uint

	fmt.Scan(&userName)
	fmt.Scan(&userTickets)

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
