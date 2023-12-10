package main

import (
	"fmt"
	"math"
	hoofd "mywebserver/morestrings"
	functions "mywebserver/tourpackage"
	"net/http"
)

var (
	ToBe           bool   = false
	MaxInt         uint64 = 1<<64 - 1
	loopTestAmount int    = 10
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Serve favicon.ico if the request is for it
	if r.URL.Path == "/favicon.ico" {
		http.ServeFile(w, r, "favicon.ico")
		return
	}
	randomprints(w, r)

}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func randomprints(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, hoofd.ReverseRunes("!oG ,niboR olleH"))
	fmt.Fprintf(w, "Now you have %g problems.\n", math.Sqrt(7))
	fmt.Fprintln(w, functions.Add(1, 2))
	fmt.Fprintf(w, "Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Println("Handling request:", r.URL.Path, "Method:", r.Method)
	fmt.Fprintf(w, "the sum = %v\n", functions.Forloopstest(&loopTestAmount))
	fmt.Fprintln(w, functions.Sqrt(2))
}
