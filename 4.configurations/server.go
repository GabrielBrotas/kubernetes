package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":4000", nil)
}

func Hello(w http.ResponseWriter, req *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	
	fmt.Fprintf(w, "Hello, I'm %s. I'm %s years old", name, age)
}