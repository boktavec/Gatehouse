package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)


type AuditEvent struct {
	ID        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	User      string    `json:"user"`
	Action    string    `json:"action"`
	Resource  string    `json:"resource"`
	IP        string    `json:"ip"`
}


func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}

var events []AuditEvent

func createEvents(w http.ResponseWriter, r *http.Request) {
	var req AuditEvent
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	// if err := json.NewDecoder(req.User).Decode()
}

func main() {
	http.HandleFunc("/health", health)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
