package main

import (
	"fmt"
	"log"
	"time"
)

var s = "Twenty One"

type User struct {
	Name string
	Phone int
	Age int
	Gender string
	Birth time.Time
}

func typesAndStructs (){
	// fmt.Println("No is",s);
	// var no2 = "Fifteen"
	// fmt.Println("Second No. is",no2)

	// var greetings,greetings2 = greet("Hello")
	// fmt.Println(greetings,greetings2)

	var user2 User

	user2 = User{
		Name : "Tanisha Tomar",
		Phone : 9191919192,
		Age : 18,
		Gender: "female",
	}
	
	user := User{
		Name : "Prince Tomar",
		Phone : 9191919191,
		Age : 21,
		Gender: "male",
	}
	fmt.Println(user)
	fmt.Println(user2)

}

func greet(s string ) (string,string){
	log.Println(s)
	return s, "Prince"
}