package main

import (
	"fmt"
	"io"
	"net/http"
)

func countHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		fmt.Fprint(w, "Send a POST request with text to count words")

	case http.MethodPost:

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, len(body))

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func main() {

	http.HandleFunc("/count", countHandler)

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", nil)
}