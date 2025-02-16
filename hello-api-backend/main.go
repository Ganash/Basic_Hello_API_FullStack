package main

import (

	"encoding/json"
	"log"
	"net/http"
)

type Message struct {

	Text string `json:"message"`

}

func helloHandler (w http.ResponseWriter, r *http.Request) {


	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")


	if r.Method == "OPTIONS" {


		w.WriteHeader(http.StatusOK)
		return

	}



	w.Header().Set("Content-Type", "application/json")
	
	response := Message{Text: "Hello From Go!"}

	json.NewEncoder(w).Encode(response)

}

func main () {


	http.HandleFunc("/api/hello", helloHandler)

	log.Println ("Server Started On: 8080")

	log.Fatal (http.ListenAndServe(":8080", nil))

}

