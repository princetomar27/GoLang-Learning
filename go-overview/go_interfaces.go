package main

import "fmt"

type Animal interface {
	Speak() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		"ruby",
		"Pit Bull Terrier",
	}

	PrintAnimalInfo(&dog)

	gorilla := Gorilla{
		"Goe",
		"Dark grey",
		34,
	}

	PrintAnimalInfo(&gorilla)
}

func PrintAnimalInfo(ani Animal) {
	fmt.Println("Animal speaks : ", ani.Speak(), " and has ", ani.NumberOfLegs(), " legs.")
}

func (d *Dog) Speak() string {
	return "Woof!"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorilla) Speak() string {
	return "Grunts and grumbles"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}
