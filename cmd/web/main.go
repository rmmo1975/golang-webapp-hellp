package main

import (
	"fmt"
	"net/http"

	"example.com/pkg/handlers"
)

const portNumber string = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting webApp on port %s", portNumber))

	_ = http.ListenAndServe(":8080", nil)
}
