package helpers

import (
	"math/rand"
	"time"
)

type HelperType struct {
	Typename string
	TypeNum  int
}

func RandomNumber(n int) int {

	rand.Seed(time.Now().UnixNano())

	value := rand.Intn(n)

	return value
}
