package main

import (
	"fmt"
	"log"
)

func main() {

	cat := "cat"

	if cat == "cat" {
		fmt.Println("Cat is cat")
	} else {
		fmt.Println("Cat is not cat.")
	}

	myVar := "cat"
	switch myVar {
	case "cat":
		log.Println("Cat is set to cat")
	case "dog":
		log.Println("Cat is set to dog")
	case "lion":
		log.Println("Cat is set to lion")
	case "fish":
		log.Println("Cat is set to fish")
	default:
		log.Println("Cat is something else ")
	}
}
