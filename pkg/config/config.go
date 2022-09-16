package config

import "html/template"

//Appconfig holds the application configuration
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
