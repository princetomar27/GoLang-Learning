package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	userMap := make(map[string]Person)

	userMap["p"] = Person{"Prince", 20}
	userMap["t"] = Person{"Tom", 18}
	userMap["g"] = Person{"Geer", 19}

	fmt.Println(userMap["p"].Name)
	fmt.Println(userMap["t"])
	fmt.Println(userMap["g"])
	//myMap := make(map[string]string)

}
