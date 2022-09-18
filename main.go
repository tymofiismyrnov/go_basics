package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName string = "Go conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}


var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidEmail && isValidName && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := FirstNames()
			fmt.Printf("The first names of bookgins are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Your email is not valid")
			}
			if !isValidTicketNumber {
				fmt.Printf("The number of tickets you've entered less than 0 or more than we have left.\nTickets remaining: %v\n", remainingTickets)
			}
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking-application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func FirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of ticket you would like to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{

		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(30 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	println("#############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	println("#############")
	wg.Done()
}
