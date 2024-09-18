package main

import (
	"fmt"
	"github.com/princetomar27/basic_web_application/pkg/handlers"
	"net/http"
)

var postNum = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Server is running on %s", postNum))
	_ = http.ListenAndServe(postNum, nil) // Start listening to the server
}
