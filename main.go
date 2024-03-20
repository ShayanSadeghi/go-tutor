package main

import "fmt"

func main() {
	const conferenceName = "GO conference"
	const totalTickets = 50
	var remainingTickets = 50

	fmt.Printf("Hello and welcome to %v. Here you can buy tickets\n", conferenceName)
	fmt.Printf("There is total %v tickets and %v are still available\n", totalTickets, remainingTickets)
	fmt.Println("Enjoy the conference!")
}
