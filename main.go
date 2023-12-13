package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/takehome/ConsumeFromKafka"
	"example.com/takehome/handlePostRequest"
)

func main() {
	go ConsumeFromKafka.ConsumeFromKafka()
	http.HandleFunc("/sendData", handlePostRequest.HandlePostRequest)
	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
