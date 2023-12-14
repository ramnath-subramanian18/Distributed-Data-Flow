package main

import (
	"fmt"
	"net/http"
	"time"
)

// sleep operation
func heavyTask() {

	time.Sleep(10 * time.Second)

	fmt.Println("Compute-heavy task completed")
}

func handleRequests(w http.ResponseWriter, r *http.Request) {
	//goroutine to call the heavy task
	go heavyTask()

	fmt.Fprintf(w, "Request received. Performing compute-heavy task...\n")
}

func main() {
	http.HandleFunc("/heavytask", handleRequests)

	fmt.Println("Server started at :8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
