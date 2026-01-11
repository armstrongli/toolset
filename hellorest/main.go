package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Handler that returns "OK" for all requests
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "OK")
	})

	// Start server on port 8080
	fmt.Println("Server starting on :8080")
	fmt.Println("Try: curl http://localhost:8080")
	fmt.Println("Or: curl http://localhost:8080/any/path")
	
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}