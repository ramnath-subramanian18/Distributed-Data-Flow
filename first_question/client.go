package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func sendRequests() {
	client := &http.Client{}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			req, err := http.NewRequest("GET", "http://localhost:8084/heavytask", nil)
			if err != nil {
				fmt.Println("Request creation failed:", err)
				return
			}

			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Request failed:", err)
				return
			}
			defer resp.Body.Close()

			fmt.Println("Request completed")
		}()
	}
	wg.Wait()
}

func main() {
	go func() {
		for {
			sendRequests()
			time.Sleep(5 * time.Second) // Adjust the interval between requests
		}
	}()

	fmt.Println("Client sending requests...")
	// Keep the client running
	select {}
}
