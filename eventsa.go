package main

import (
	"encoding/json"
	"net/http"
)

func getAllEventsA(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}
