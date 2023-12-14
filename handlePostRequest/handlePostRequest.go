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

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error encoding data", http.StatusInternalServerError)
		return
	}
	//get the data from the post request and send to kafka
	//goroutine-1.
	go sendDataToKafka.SendDataToKafka(jsonData, "Data_topic")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Data sent to Kafka!")
}
