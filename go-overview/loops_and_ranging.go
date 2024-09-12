package main

import (
	"fmt"
	"log"
)

func main() {
	// Loop through even numbers from 0 to 8
	for i := 0; i < 10; i += 2 {
		fmt.Print(i, " ")
	}
	fmt.Println() // Newline after printing numbers

	// Map of car brands and models
	cars := map[string]string{
		"Volvo":    "XC90",
		"Suzuki":   "Wagon R",
		"Mahindra": "Scorpio",
		"Toyota":   "Corolla",
		"Honda":    "Civic",
	}

	// Print car models using range
	for _, car := range cars {
		fmt.Print(car, ", ")
	}
	fmt.Println()

	// Print car brands and their respective models
	for brand, model := range cars {
		fmt.Printf("%s makes %s\n", brand, model)
	}

	// Define a User struct
	type User struct {
		Name  string
		Age   int
		Email string
	}

	// Slice of User structs
	users := []User{
		{"John", 25, "john@gmail.com"},
		{"Wick", 18, "wick@gmail.com"},
		{"Johnny", 32, "johnny@gmail.com"},
		{"Wicky", 29, "wicky@gmail.com"},
	}

	// Loop through users and log their details
	for _, user := range users {
		log.Printf("%s is %d years old. Email: %s\n", user.Name, user.Age, user.Email)
	}
}
