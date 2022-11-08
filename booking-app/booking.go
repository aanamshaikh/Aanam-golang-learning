package main

import (
	"fmt"
	"strings"
)

// printing into console
func main() {

	appName := "Booking App"
	// or this var appName = "Booking App"

	// this value will not change throughout the program
	const totalTickets int = 50
	var availableTickets int = 50
	var bookings []string //i.e SLICE

	fmt.Println("Welome to our " + appName)
	fmt.Println("Hurry up only", availableTickets, " tickets are remaining !!!")

	for {
		var firstName string
		var lastName string
		var email string

		var userTickets int
		// userName="Aanam"
		// userTickets=1

		fmt.Print("Please Enter your First Name: ")
		fmt.Scan(&firstName)
		fmt.Print("Please Enter your Last Name: ")
		fmt.Scan(&lastName)

		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2

		if !isValidName {
			fmt.Println("Your name should atleast have 2 Charaters")
			continue
		}

		fmt.Print("Please Enter your Email Address: ")
		fmt.Scan(&email)

		var isValidEmail = strings.Contains(email, "@")
		if !isValidEmail {
			fmt.Println("Please Enter a Valid Email Address")
			continue
		}

		fmt.Print("Please Enter number of tickets you wish to purchase: ")
		fmt.Scan(&userTickets)

		if userTickets < availableTickets {
			availableTickets = availableTickets - userTickets

			// var bookings = [50]string{} //or
			// bookings[0] = firstName + " " + lastName

			//SLICE

			// var bookings []string //i.e SLICE
			bookings = append(bookings, firstName+" "+lastName)
			firstNames := []string{}
			for _, booking := range bookings {
				var names []string = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			// fmt.Printf("First Value of Array %v \n",bookings[0])
			// fmt.Printf("Length of Array %v \n",len(bookings))

			fmt.Printf("Thankyou %v %v for booking %v tickets.\nThe invoice will be sent to you via Email at %v \n", firstName, lastName, userTickets, email)
			fmt.Println("Hurry up only", availableTickets, "tickets are remaining !!!")
			fmt.Printf("List of Bookings :%v\n", firstNames)

		} else if userTickets == availableTickets {
			availableTickets = availableTickets - userTickets
			fmt.Printf("Looks Like you're our last customer :)) ,Thankyou\n Invoice will be sent to you via Email at %v\n", email)
			var noTicketsRemaining bool = availableTickets == 0
			if noTicketsRemaining {
				fmt.Print("Sorry,all the tickets are sold out :)\n Press Q if you wish to quit ")

				var input string
				fmt.Scan(&input)
				if input == "Q" {
					fmt.Println("Thankyou!!!")
					break
				}
			}
		} else {
			fmt.Printf("We only have %v tickets remaining\n", availableTickets)
			continue
		}

	}

}