package main

import "net/http"

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":4000", nil)
}

func Hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello from go http k8s server. v3"))
}