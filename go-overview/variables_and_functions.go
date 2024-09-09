package main

import "fmt"

func main(){
	fmt.Println("Hello World !")

	var name string
	name = "Prince"

	var i int
	i = 21


	fmt.Println("This world is cruel",name)
	fmt.Println("You are",i,"years old.")

	whatWasSaid,age,measure := saySomething()
	
	fmt.Println(whatWasSaid,age,measure)
}

func saySomething() (string,int,string){
	return "Good Bye!",10,"years old"
}