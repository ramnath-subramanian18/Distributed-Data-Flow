package handlePostRequest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/takehome/sendDataToKafka"
)

func HandlePostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data map[string]interface{}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Convert data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error encoding data", http.StatusInternalServerError)
		return
	}

	// Send data to Kafka using goroutine
	go sendDataToKafka.SendDataToKafka(jsonData, "your_topic")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Data sent to Kafka!")
}
