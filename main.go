package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to golang V2",
		Description: "Come join us for a chance to learn how golang works",
	},
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data error")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newEvent)
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventId := mux.Vars(r)["id"]
	for _, singleEvent := range events {
		if singleEvent.ID == eventId {
			json.NewEncoder(w).Encode(singleEvent)
			return
		}
	}
	fmt.Fprintf(w, `{"message": "not found"}`)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/event", createEvent).Methods(http.MethodPost)
	router.HandleFunc("/event/{id}", getOneEvent).Methods(http.MethodGet)
	router.HandleFunc("/event", getAllEvents).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8888", router))
}
