package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PersonGreeting struct {
	Person   Person `json:"person"`
	Greeting string `json:"greeting"`
}

func main() {
	//defining channel for results of goroutines.
	var c chan string
	c = make(chan string, 2)

	person := getPerson()
	go postGreeting(person, c)
	go postGreeting(person, c)

	for i := range c {
		//printing goroutine status from channel.
		fmt.Println(i)
	}
}

func getPerson() Person {
	bodyBytes, err := CallPersonAPI("GET", nil)
	if err != nil {
		//If error occurs while getting person no need to continue on this code. So I used panic instead of log.
		panic("An error occured while getting Person from API: " + err.Error())
	}
	var person Person
	json.Unmarshal(bodyBytes, &person)
	return person
}

func postGreeting(person Person, c chan string) {
	greeting := new(PersonGreeting)
	greeting.Person = person
	greeting.Greeting = fmt.Sprintf("Hello %s (%d)", person.Name, person.Age)
	jsonData, err := json.Marshal(greeting)
	_, err = CallPersonAPI("POST", jsonData)
	if err != nil {
		fmt.Println("An error occured while posting Person to API: ", err.Error())
		c <- err.Error()
	} else {
		c <- "Success"
	}
}

func CallPersonAPI(method string, jsonData []byte) ([]byte, error) {
	fmt.Println("Calling Person API...")
	client := &http.Client{}
	req, err := http.NewRequest(method, "http://localhost:8080/person", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("An error occured while creating request for API", err.Error())
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("An error occured on API call: ", err.Error())
		return nil, err
	} else if resp.StatusCode == http.StatusOK {
		fmt.Printf("Request has been successfully completed: %s %s://%s%s\n", req.Method, req.URL.Scheme, req.URL.Host, req.URL.Path)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("An error occured while reading body of the request: ", err.Error())
		return nil, err
	}
	return bodyBytes, err
}
