package main

import (
	"fmt"
	"log"
)

func main() {
	for i := 0; i < 10; i += 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()

	//animals := []string{"cat", "dog", "lion", "tiger", "none"}

	cars := make(map[string]string)
	cars["volvo"] = "Volvo"
	cars["suzuki"] = "Wagon r"
	cars["Mahindra"] = "Scorpion"
	cars["Toyota"] = "Corolla"
	cars["Honda"] = "Civic"

	//for i, animal := range animals {
	//	fmt.Println(i+1, " ", animal)
	//}

	for _, car := range cars {
		fmt.Print(car, ", ")
	}
	for brand, car := range cars {
		fmt.Println(brand, " makes ", car)
	}

	type User struct {
		Name  string
		Age   int
		Email string
	}

	var users = []User{}
	users = append(users, User{"John", 25, "john@gmail.com"})
	users = append(users, User{"Wick", 18, "wick@gmail.com"})
	users = append(users, User{"Johnny", 32, "johnny@gmail.com"})
	users = append(users, User{"Wicky", 29, "wicky@gmail.com"})

	for _, currentUser := range users {
		log.Println(currentUser.Name, " is ", currentUser.Age, " years old.\n", "Email is ", currentUser.Email, "\n")
	}
}
