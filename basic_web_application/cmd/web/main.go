package main

import (
	"fmt"
	"github.com/princetomar27/basic_web_application/pkg/config"
	"github.com/princetomar27/basic_web_application/pkg/handlers"
	"github.com/princetomar27/basic_web_application/pkg/render"
	"log"
	"net/http"
)

var postNum = ":8080"

func main() {

	var app config.AppConfig
	tc, err := render.CreateCompTemplateCache()
	if err != nil {
		log.Fatal("Got error while creating template cache : ", err)
	}

	app.TemplateCache = tc
	app.UseCache = false
	// set things up for our handlers, by creating a repo variable which calls the function handlers.NewHandlers

	repo := handlers.NewRepo(&app)

	// pass repo back to the handlers
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Server is running on %s", postNum))
	_ = http.ListenAndServe(postNum, nil) // Start listening to the server
}
