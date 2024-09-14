package main

import (
	"fmt"
	"github.com/princetomar27/myfirstpackage/helpers"
)

func main() {
	fmt.Println("Hello Prince, Hey !")

	var mySomeTime helpers.HelperType

	mySomeTime.Typename = "Prince"
	mySomeTime.TypeNum = 21

	fmt.Println(mySomeTime)
}
