package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/greet/", Greet)

	http.ListenAndServe(":8081", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func Greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/greet/"):]
	fmt.Fprintf(w, "Hello %s\n", name)
}
