package config

import "html/template"

// THIS PACKAGE IS IMPORTED IN OTHER PARTS OF THE APPLICATION
// BUT THIS ISN'T IMPORTING ANYTHING

// AppConfig holds the Application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
