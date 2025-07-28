package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Dockerized Go App!")
	})

	fmt.Println("Server running at :8080")
	http.ListenAndServe(":8080", nil)
}
