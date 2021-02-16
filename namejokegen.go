package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type NameResponse struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

// { "type": "success", "value": { "id": 181, "joke": "John Doe's OSI network model has only one layer - Physical.", "categories": ["nerdy"] } }
type JokeResponse struct {
	Value struct {
		Joke string `json:"joke"`
	} `json:"value`
}


func main()  {
	// puts the response on the server
	http.HandleFunc("/", genServer)
	http.ListenAndServe(":5000", nil)
}

func genServer(w http.ResponseWriter, r *http.Request)  {
	// Loading in the name
	nResponse, err := http.Get("https://names.mcquay.me/api/v0/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	nData, err := ioutil.ReadAll(nResponse.Body)

	if err != nil {
		log.Fatal(err)
	}
	
	var name NameResponse
	json.Unmarshal(nData, &name)

	
	// Loading in the joke
	jokeLink := fmt.Sprintf("http://api.icndb.com/jokes/random?firstName=%s&lastName=%s&limitTo=\\[nerdy\\]", name.FirstName, name.LastName)
	jResponse, err := http.Get(jokeLink)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	jData, err := ioutil.ReadAll(jResponse.Body)

	if err != nil {
		log.Fatal(err)
	}

	var j JokeResponse
	json.Unmarshal(jData, &j)
	fmt.Fprintf(w, "%s", j.Value.Joke)
}