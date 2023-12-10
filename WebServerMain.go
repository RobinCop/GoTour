package main

import (
	"fmt"
	"math"
	hoofd "mywebserver/morestrings"
	functions "mywebserver/tourpackage"
	"net/http"
)

var (
	ToBe   bool   = false
	MaxInt uint64 = 1<<64 - 1
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Serve favicon.ico if the request is for it
	if r.URL.Path == "/favicon.ico" {
		http.ServeFile(w, r, "favicon.ico")
		return
	}
	r.URL.User.Password()

	fmt.Fprintln(w, hoofd.ReverseRunes("!oG ,niboR olleH"))
	fmt.Fprintf(w, "Now you have %g problems.\n", math.Sqrt(7))
	fmt.Fprintln(w, functions.Add(1, 2))
	fmt.Fprintf(w, "Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Println("lolollo")
	fmt.Println("Handling request:", r.URL.Path, "Method:", r.Method)

}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
