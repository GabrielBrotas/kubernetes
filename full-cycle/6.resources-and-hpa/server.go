package main

import (
	"fmt"
	"net/http"
	"time"
)

var startedAt = time.Now() 

func main() {
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":4000", nil)
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello,")
}

func Healthz(w http.ResponseWriter, req *http.Request) {

	// how long the application is running
	duration := time.Since(startedAt) 

	// seconds_to_throw_error := 60.0
	seconds_to_application_start := 10.0

	// if duration.Seconds() > seconds_to_throw_error || duration.Seconds() < seconds_to_application_start {
	if duration.Seconds() < seconds_to_application_start {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration)))
		return;
	}

	w.WriteHeader(200)
	w.Write([]byte("ok"))
}
