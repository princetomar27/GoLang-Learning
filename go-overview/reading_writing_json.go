package main

import (
	"encoding/json"
	"log"
)

type Persons struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	myJson := `
[
	{
		"first_name":"Clark",
		"last_name":"Kent",
		"age":37,
		"has_dog":true
	},
	{
		"first_name":"Bruce",
		"last_name":"Wayne",
		"age":39,
		"has_dog":false
	}
]`

	var unmarshalled []Persons

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling JSON: ", err)
	}

	log.Printf("Unmarshalled: %v", unmarshalled)

	// write json from a struct
	var myPeople []Persons

	var p1 Persons
	p1.FirstName = "Prince"
	p1.LastName = "Tomar"
	p1.Age = 21
	p1.HasDog = true

	myPeople = append(myPeople, p1)

	var p3 Persons
	p3.FirstName = "A"
	p3.LastName = "B"
	p3.Age = 19
	p3.HasDog = true

	myPeople = append(myPeople, p3)

	newJson, err := json.MarshalIndent(myPeople, "", "	")

	if err != nil {
		log.Println("Error marshalling JSON: ", err)
	}
	log.Println(string(newJson))
}
