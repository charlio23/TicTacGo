package main

import (
	//"encoding/json"
	"fmt"
	//"github.com/gorilla/mux"
	//"log"
	//"net/http"
	//"os"
)

/*
type Val struct {
	Value string `json:"value,omitempty"`
}

type Vals []Val

func parseBoard(value string) [3][3]string {
	var board [3][3]string
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = value[3*i+j]
		}
	}
	return board
}

func encodeBoard(board [3][3]string) string {
	value := ""
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			value += board[i][j]
		}
	}
	return value
}

func index(w http.ResponseWriter, r *http.Request) {
	vals := Vals{
		Val{Value: "Hello world"},
	}
	json.NewEncoder(w).Encode(vals)
}
*/
func main() {
	/*
		result := parseBoard("---------")
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				fmt.Print(" " + board[i][j])
			}
			fmt.Println()
		}
		port := os.Getenv("PORT")
		router := mux.NewRouter()
		router.HandleFunc("/nextMove", index).Methods("GET")
		log.Fatal(http.ListenAndServe(":"+port, router))*/
	var origBoard [9]string = [9]string{"0", "1", "X", "O", "O", "5", "O", "X", "X"}
	fmt.Println(callAi(origBoard))
}
