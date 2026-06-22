package main

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func main() {
	http.HandleFunc("/ping", pingHandler)

	fmt.Println("Server is starting on port 8080...")


	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}