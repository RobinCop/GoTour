package main

import (
	"fmt"
	"math"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go Web Server!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on http://localhost:8080")
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println("But a bitch aint one")
	http.ListenAndServe(":8080", nil)
}
