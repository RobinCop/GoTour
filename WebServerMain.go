package main

import (
	"fmt"
	"math"
	"mywebserver/morestrings"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go Web Server!")
	fmt.Fprintf(w, "Now you have %g problems.\n", math.Sqrt(7))
	fmt.Fprintln(w, "But a bitch aint one")
	fmt.Fprintln(w, morestrings.ReverseRunes("!oG ,olleH"))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
