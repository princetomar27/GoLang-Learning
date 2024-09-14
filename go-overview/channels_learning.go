package main

import (
	"github.com/princetomar27/myfirstpackage/helpers"
	"log"
)

const numberPool = 100

func CalculateChanValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numberPool)
	intChan <- randomNumber
}

func main() {

	intChan := make(chan int)

	defer close(intChan)

	go CalculateChanValue(intChan)

	num := <-intChan

	log.Println(num)

}
