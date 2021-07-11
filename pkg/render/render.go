package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"example.com/pkg/config"
	"example.com/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func SetAppConfig(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, templatePage string, templateData *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[templatePage]
	if !ok {
		log.Fatal("template not found from cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, templateData)

	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	rootTemplatesDir := os.Getenv("TEMPLATES_PATH")

	pages, err := filepath.Glob(rootTemplatesDir + "/templates/*-page.html")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		// just to inform that the current page loaded
		fmt.Println("current page loaded:", page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		layouts, err := filepath.Glob(rootTemplatesDir + "/templates/*-layout.html")

		if err != nil {
			return myCache, err
		}

		if len(layouts) > 0 {
			ts, err := ts.ParseGlob(rootTemplatesDir + "/templates/*-layout.html")
			if err != nil {
				return myCache, err
			}
			myCache[name] = ts
		}

	}
	return myCache, nil
}
