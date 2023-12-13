package main

import (
	"fmt"
	"net/http"
	"time"
)

func heavyTask() {
	// Simulate a compute-heavy task by sleeping for some time
	time.Sleep(10 * time.Second) // Change the duration as needed

	fmt.Println("Compute-heavy task completed")
}

func handleRequests(w http.ResponseWriter, r *http.Request) {
	go heavyTask()

	fmt.Fprintf(w, "Request received. Performing compute-heavy task...\n")
}

func main() {
	http.HandleFunc("/heavytask", handleRequests)

	fmt.Println("Server started at :8084")
	if err := http.ListenAndServe(":8084", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
