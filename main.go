package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

type helloWorldResponse struct {
	Message string `json:"message"`
	Author  string `json:"-"`
	Date    string `json:",omitempty"`
	Id      int    `json:"id, string"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func main() {
	port := 8989

	http.HandleFunc("/", helloHandler)

	log.Printf("Server starting on port %v \n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic("Failed to read Response Body!")
	}
	response := helloWorldResponse{Message: "Hello World"}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
