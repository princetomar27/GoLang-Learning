package main

import "fmt"

func pointersMain() {

	var myColor string 
	myColor = "black"

	fmt.Println("I like color :",myColor)
	fmt.Println("But after the update")
	updateViaPointer(&myColor)
	fmt.Println("I like color :",myColor)
}

func updateViaPointer(color *string){
	newColor := "Red"
	*color = newColor
}