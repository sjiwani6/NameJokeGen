package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"html"
)

/* 
	NameResponse is a struct to hold the relevant values from the https://names.mcquay.me/api/v0/ API
*/
type NameResponse struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}
/*
	JokeResponse is a struct to hold the needed joke from the Chuck Norris Joke API
*/
type JokeResponse struct {
	Value struct {
		Joke string `json:"joke"`
	} `json:"value"`
}


func main()  {
	// puts the joke on the web server hosted on http://localhost:5000
	http.HandleFunc("/", GenServer)
	http.ListenAndServe(":5000", nil)
}

/* 
	GenServer pulls the first and last name from the first API, then inserts the first and last
	name into the URL for the joke's API. Then, it will get a response from the joke API which
	will be used to construct an HTTP response.
*/
func GenServer(w http.ResponseWriter, r *http.Request)  {
	// HTTP GET call to the API
	nResponse, err := http.Get("https://names.mcquay.me/api/v0/")

	// Error checking
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// Turning the response into a NameResponse object
	nData, err := io.ReadAll(nResponse.Body)
	if err != nil {
		log.Fatal(err)
	}
	var name NameResponse
	json.Unmarshal(nData, &name)

	
	// HTTP GET call to the joke API
	jokeLink := fmt.Sprintf("http://api.icndb.com/jokes/random?firstName=%s&lastName=%s&limitTo=\\[nerdy\\]", name.FirstName, name.LastName)
	jResponse, err := http.Get(jokeLink)

	// Error handling
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// Turning the response into a JokeResponse object
	jData, err := io.ReadAll(jResponse.Body)
	if err != nil {
		log.Fatal(err)
	}
	var j JokeResponse
	json.Unmarshal(jData, &j)

	// This line will unescape any entities in the Joke response
	newJoke := html.UnescapeString(j.Value.Joke)
	fmt.Fprintf(w, "%s", newJoke)
}