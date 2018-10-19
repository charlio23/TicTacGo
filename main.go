package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Val struct {
	Value string `json:"value,omitempty"`
}

type Vals []Val

func index(w http.ResponseWriter, r *http.Request) {
	vals := Vals{
		Val{Value: "Hello world"},
	}
	json.NewEncoder(w).Encode(vals)
}

func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
