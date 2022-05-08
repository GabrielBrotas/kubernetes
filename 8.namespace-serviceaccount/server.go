package main

import (
	"fmt"
	"net/http"
)


func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":4000", nil)
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello,")
}

