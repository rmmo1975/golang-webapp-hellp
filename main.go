package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber string = ":8080"

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home-page.html")
}

func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about-page.html")
}

func renderTemplate(w http.ResponseWriter, templatePage string) {
	parsedTemplatePage, _ := template.ParseFiles("./templates/" + templatePage)

	err := parsedTemplatePage.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsin template:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	fmt.Println(fmt.Sprintf("Starting webApp on port %s", portNumber))

	_ = http.ListenAndServe(":8080", nil)
}
