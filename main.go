package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName string = "Go conference"
var RemainingTickets uint = 50

// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	// var reservations = [50]string{}

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, RemainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		noTicketsRemaining := RemainingTickets == 0
		if noTicketsRemaining {
			fmt.Println(" Our conference is booked out. Come back next year")
		}
	} else {
		if !isValidName {
			fmt.Println("First name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sing")
		}

		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, RemainingTickets)

	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
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

	// ask user for their personal data
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	RemainingTickets = RemainingTickets - userTickets
	// create a map for a user
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	//reservations[0] = firstName + " " + lastName

	// fmt.Printf("The whole array: %v\n", reservations)
	// fmt.Printf("The first value: %v\n", reservations[0])
	// fmt.Printf("Array type: %T\n", reservations)
	// fmt.Printf("Array length: %v\n", len(reservations))

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Slice type: %T\n", bookings)
	// fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", RemainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println(strings.Repeat("#", 10))
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println(strings.Repeat("#", 10))
	wg.Done()
}
