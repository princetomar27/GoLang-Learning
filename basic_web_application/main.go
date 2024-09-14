package main

import (
	"fmt"
	"net/http"
)

var postNum = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "this is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {

	sum := addValues(3, 4)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and sum is %d", sum))
}

func addValues(x, y int) int {
	return x + y
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println(fmt.Sprintf("Server is running on %s", postNum))
	_ = http.ListenAndServe(postNum, nil) // Start listening to the server
}
