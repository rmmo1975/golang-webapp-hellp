package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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

	app.UseCache, _ = strconv.ParseBool(os.Getenv("USE_CACHE"))
	app.TemplateCache = tc

	// how to get rid of this way of handlers start? using a kind of singleton.
	repo := handlers.CreateRepo(&app)
	handlers.StartHandlers(repo)

	render.SetAppConfig(&app)

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Println(fmt.Sprintf("Starting webApp on port %s", portNumber))
	serve.ListenAndServe()
}
