package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PersonGreeting struct {
	Person   Person `json:"person"`
	Greeting string `json:"greeting"`
}

func getPerson(w http.ResponseWriter, _ *http.Request) {
	person := Person{Name: "William", Age: 42}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(person)
}

func createPerson(w http.ResponseWriter, req *http.Request) {
	time.Sleep(5 * time.Second) // Fake long request

	var personGreeting PersonGreeting
	if err := json.NewDecoder(req.Body).Decode(&personGreeting); err != nil {
		_, _ = fmt.Fprintf(w, "Error parsing JSON: "+err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	log.Printf("Person: %v\nRegistered greeting: %s\n",
		personGreeting.Person, personGreeting.Greeting)
}

func handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		getPerson(w, req)
	} else if req.Method == http.MethodPost {
		createPerson(w, req)
	} else {
		_, _ = fmt.Fprintf(w, "Method is not supported")
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func main() {
	http.HandleFunc("/person", handler)

	log.Println("\nRegistered endpoints:\n\t* GET /person\n\t* POST /person\nRunning on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
