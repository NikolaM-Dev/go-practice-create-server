package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func HandlePostUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user User

	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid data")

		return
	}

	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var metaData MetaData

	err := decoder.Decode(&metaData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid data")

		return
	}

	fmt.Fprintf(w, "Payload %v\n", metaData)
}
