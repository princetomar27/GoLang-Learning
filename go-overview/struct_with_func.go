package main

import "fmt"

type myStruct struct {
	Name string
	Age  int
}

func (person *myStruct) printNameAndAge() (string, int) {
	return "Your name is : " + person.Name + "\nAnd your age is", person.Age
}

func (person *myStruct) decreaseAge() int {
	return person.Age - 10
}

func (person *myStruct) increaseAge() int {
	return person.Age + 10
}

func structWithFunc() {

	var prince myStruct

	prince.Name = "Prince"
	prince.Age = 21

	fmt.Println(prince)
	fmt.Println(prince.printNameAndAge())
	fmt.Println(prince.decreaseAge())
	fmt.Println(prince.increaseAge())
	fmt.Println("Current Age is : ", prince.Age)
}

func main() {
	structWithFunc()
}
