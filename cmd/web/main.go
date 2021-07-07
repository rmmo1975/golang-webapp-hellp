package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/pkg/config"
	"example.com/pkg/handlers"
	"example.com/pkg/render"
)

const portNumber string = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app.UseCache = false
	app.TemplateCache = tc

	repo := handlers.CreateRepo(&app)
	handlers.StartHandlers(repo)

	render.SetAppConfig(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting webApp on port %s", portNumber))

	_ = http.ListenAndServe(":8080", nil)
}
