package handlers

import (
	"errors"
	"fmt"
	"github.com/princetomar27/basic_web_application/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//_, _ = fmt.Fprintf(w, "this is the home page")
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {

	//sum := addValues(3, 4)
	//_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and sum is %d", sum))
	render.RenderTemplate(w, "about.page.tmpl")
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
