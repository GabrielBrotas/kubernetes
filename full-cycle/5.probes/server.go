package main

import (
	"fmt"
	"net/http"
	"time"
)

var startedAt = time.Now() 

func main() {
	http.HandleFunc("/healthz", Healthz_StartupProbe)
	http.HandleFunc("/healthz1", Healthz_Liveness)
	http.HandleFunc("/healthz2", Healthz_Readiness)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":4000", nil)
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello,")
}

func Healthz_StartupProbe(w http.ResponseWriter, req *http.Request) {

	// how long the application is running
	duration := time.Since(startedAt) 

	seconds_to_throw_error := 30.0
	seconds_to_application_start := 10.0

	if duration.Seconds() > seconds_to_throw_error || duration.Seconds() < seconds_to_application_start {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration)))
		return;
	}

	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func Healthz_Liveness(w http.ResponseWriter, req *http.Request) {

	// how long the application is running
	duration := time.Since(startedAt) 

	seconds_to_throw_error := 50.0
	if duration.Seconds() > seconds_to_throw_error {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration)))
		return;
	}

	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func Healthz_Readiness(w http.ResponseWriter, req *http.Request) {

	// how long the application is running
	duration := time.Since(startedAt) 

	seconds_to_application_start := 10.0
	// coldtime until application start so we can test the readiness
	// we don't wanna traffic until this coldtime pass
	if duration.Seconds() < seconds_to_application_start {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration)))
		return;
	}

	w.WriteHeader(200)
	w.Write([]byte("ok"))
}