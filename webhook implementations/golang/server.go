package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//APIAIRequest : Incoming request format from APIAI
type APIAIRequest struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Result    struct {
		Parameters map[string]string `json:"parameters"`
		Contexts   []interface{}     `json:"contexts"`
		Metadata   struct {
			IntentID                  string `json:"intentId"`
			WebhookUsed               string `json:"webhookUsed"`
			WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed"`
			IntentName                string `json:"intentName"`
		} `json:"metadata"`
		Score float32 `json:"score"`
	} `json:"result"`
	Status struct {
		Code      int    `json:"code"`
		ErrorType string `json:"errorType"`
	} `json:"status"`
	SessionID       string      `json:"sessionId"`
	OriginalRequest interface{} `json:"originalRequest"`
}

//APIAIMessage : Response Message Structure
type APIAIMessage struct {
	Speech      string `json:"speech"`
	DisplayText string `json:"displayText"`
	Source      string `json:"source"`
}

//HelloEndpoint - HTTP Request Handler for /hello
func HelloEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello from APIAI Webhook Integration.")
}

//VersionEndpoint - HTTP Request Handler for /version
func VersionEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "APIAI Webhook Integration. Version 1.0")
}

//WebhookEndpoint - HTTP Request Handler for /webhook
func WebhookEndpoint(w http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" {
		decoder := json.NewDecoder(req.Body)

		var t APIAIRequest
		err := decoder.Decode(&t)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error in decoding the Request data", http.StatusInternalServerError)
		}

		rating := t.Result.Parameters["rating"]
		comments := t.Result.Parameters["comments"]
		resortlocation := t.Result.Parameters["resort-location"]
		fmt.Println("Received the following request parameters", rating, comments, resortlocation)
		msg := APIAIMessage{Source: "Hotel Feedback System", Speech: "Thank you for the feedback", DisplayText: "Thank you for the feedback"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msg)
	} else {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HelloEndpoint).Methods("GET")
	router.HandleFunc("/version", VersionEndpoint).Methods("GET")
	router.HandleFunc("/webhook", WebhookEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":5000", router))
}
