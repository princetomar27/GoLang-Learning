package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

var postNum = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	//_, _ = fmt.Fprintf(w, "this is the home page")
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {

	//sum := addValues(3, 4)
	//_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and sum is %d", sum))
	renderTemplate(w, "about.page.tmpl")
}

func addValues(x, y int) int {
	return x + y
}

func divideValues(x, y float32) (float32, error) {

	if y <= 0 {
		err := errors.New("y must be greater than 0")
		return 0, err
	}

	result := x / y
	return result, nil
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)

	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("100 / 10 = %f", f))
}

func renderTemplate(w http.ResponseWriter, tmpl string) {

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Server is running on %s", postNum))
	_ = http.ListenAndServe(postNum, nil) // Start listening to the server
}
