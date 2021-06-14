package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, templatePage string) {

	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	t, ok := tc[templatePage]
	if !ok {
		log.Fatal("template not found", err)
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)

	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*-page.html")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		// just to inform that the current page loaded

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		layouts, err := filepath.Glob("./templates/*-layout.html")

		if err != nil {
			return myCache, err
		}

		if len(layouts) > 0 {
			ts, err := ts.ParseGlob("./templates/*-layout.html")
			if err != nil {
				return myCache, err
			}
			myCache[name] = ts
		}

	}
	return myCache, nil
}
