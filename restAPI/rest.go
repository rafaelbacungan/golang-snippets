package restAPI

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func RestAPI() {
	// Used to instantiate the server
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Dummy database
/*
	Next, we’ll create a dummy database to store our events.

	For our event struct, we will only be using the Title and Description fields. The slice, or dummy database, should
	only hold event structs. Therefore, from the slice, we can call a new event struct, read what it says change it if
	necessary, or delete it.
*/
type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Introduction to Golang",
	},
}

// Creating an event
/*
	The data for creating a new event comes from the user's end when the user inputs data in the form of an http request
	data. When input, the request data is not in a form that is readable by humans, so to translate it into a slice, we
	use the package ioutil.

	After it has been translated into a slice we fit into an event struct by unmarshalling it. Once the slice is
	successfully created, we can append the event struct into the new events slice and show the new event with an http
	response of 201 Created Status Code.
*/
func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	fmt.Println(events)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

// Get an event
/*
	By using the GET Method, we will be able to access the endpoint for getting one event and it will look like this
	/events/{id}.

	Within Gorilla Mu, we can obtain the value to be inserted into the "id" to filter out a selected event from the
	events slice. Once an "id" that resembles the input "id" is located, its value is obtained from the events slice
	and displayed as a response to the user within the API.
*/
func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

// Get all events
/*
	In order to gather all events in the slice, you simply need to display the entire events slice.
*/
func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

// Update an event
/*
	We use a PATCH method and an endpoint of /evets/{id} to update an existing event. Again, with the help of Gorilla
	Mux, we find the value of the "id" has been located, we can change the values of the Title and Description fields
	within the event struct.

	Then, we change the value of the struct in the events slice. Next, we return the updated value for the event
	struct to the user as a response.
*/
func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &updatedEvent)

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			singleEvent.Title = updatedEvent.Title
			singleEvent.Description = updatedEvent.Description
			events[i] = singleEvent
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}
