package config

import "html/template"

// AppConfig holds the webapp config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
