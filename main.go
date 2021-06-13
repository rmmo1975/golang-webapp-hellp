package main

import (
	"fmt"
	"net/http"
)

const portNumber string = ":8080"

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	fmt.Println(fmt.Sprintf("Starting webApp on port %s", portNumber))

	_ = http.ListenAndServe(":8080", nil)
}
