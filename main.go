package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Val struct {
	Value string `json:"value,omitempty"`
}

var Vals []Val

func GetVals(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Vals)
}

func encodeBoard(board [9]string) string {
	value := ""
	for i := 0; i < 9; i++ {
		value += board[i]
	}
	return value
}

func index(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range Vals {
		if item.Value == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Val{})
}

func main() {
	//port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/nextMove", index).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
	var urls = []string{
		"http://google.com",
		"http://facebook.com",
		"http://youtube.com",
		"http://yahoo.com",
		"http://twitter.com",
		"http://live.com",
	}

	urlsJson, _ := json.Marshal(urls)
	fmt.Println(string(urlsJson))
	//var origBoard [9]string = [9]string{"0", "1", "2", "3", "4", "5", "6", "7", "8"}
	//fmt.Println(callAi(origBoard))
}
