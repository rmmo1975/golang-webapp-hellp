package main

import (
	"errors"
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

func divide(w http.ResponseWriter, r *http.Request) {
	result, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, result))
	return
}

func divideValues(x, y float32) (float32, error) {

	// check if y - divisor - is not 0 or negative
	if y <= 0 {
		return 0, errors.New("divisor is 0 or negative. not allowed.")
	}

	result := x / y

	return result, nil
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	http.HandleFunc("/divide", divide)

	fmt.Println(fmt.Sprintf("Starting webApp on port %s", portNumber))

	_ = http.ListenAndServe(":8080", nil)
}
