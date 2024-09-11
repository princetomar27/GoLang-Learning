package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type Human struct {
	Name      string
	Age       int
	behaviour string
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

	humanMap := make(map[string]Human)

	humanMap["Prince"] = Human{
		"Prince Tomar",
		21,
		"Calm",
	}

	humanMap["Tanisha"] = Human{"Tanisha", 18, "Cheerful"}

	fmt.Println(humanMap)

	var mySlice []string

	mySlice = append(mySlice, "Prince")
	mySlice = append(mySlice, "Tanisha")
	fmt.Println(mySlice)

	fmt.Println(mySlice[1])

	var myInts []int

	myInts = append(myInts, 4)
	myInts = append(myInts, 2)
	myInts = append(myInts, 3)
	myInts = append(myInts, 9)

	fmt.Println(myInts)

	for i := range myInts {
		fmt.Println(myInts[i])
	}

	sort.Ints(myInts)
	fmt.Println(myInts)
}
