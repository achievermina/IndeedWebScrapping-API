package main

import (
	//"github.com/achievermina/IndeedWebScrapping-API/scrapper"
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"encoding/json"
)

//https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da

type event struct {
	userID          string `json:"userID"`
	Title       string `json:"Title"`
}

type allEvents []event

var events = allEvents{
	{
		userID:          "1",
		Title:       "Introduction to Golang",
	},
}

var searchterm string
var allsearch []string

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	//var newSearch string
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "please type the word you are searching for")
	}

	json.Unmarshal(reqBody, &newEvent)
	//allsearch = append(allsearch, newSearch)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEvent)
}

func main() {
	//term :="python"
	//scrapper.Scrape(term)

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/event", createEvent).Methods("GET")
	log.Fatalln(http.ListenAndServe(":8080", r))

}


