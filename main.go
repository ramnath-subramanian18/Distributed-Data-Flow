package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/takehome/ConsumeFromKafka"
	"example.com/takehome/handlePostRequest"
)

func main() {
	//ConsumeFromKafka is used to fetch the data from kafka and data is stored in redis
	//goroutine-2
	go ConsumeFromKafka.ConsumeFromKafka()
	//handlepostrequest function is activated by the post request
	http.HandleFunc("/sendData", handlePostRequest.HandlePostRequest)
	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
