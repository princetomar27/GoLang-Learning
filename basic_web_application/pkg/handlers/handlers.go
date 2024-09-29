package handlers

import (
	"errors"
	"fmt"
	"github.com/princetomar27/basic_web_application/pkg/config"
	"github.com/princetomar27/basic_web_application/pkg/models"
	"github.com/princetomar27/basic_web_application/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//_, _ = fmt.Fprintf(w, "this is the home page")
	//render.RenderTemplate(w, "home.page.tmpl")
	render.RenderTemplateAdvanced(w, "home.page.tmpl", &models.TemplateData{})

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// perform some business logic here
	stringMap := make(map[string]string)
	stringMap["hello"] = "Hello, again."

	// send the data to the template
	render.RenderTemplateAdvanced(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

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
