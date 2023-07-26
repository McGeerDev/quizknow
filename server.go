package main

import (
	"fmt"
	"log"
	"net/http"

	api "svelte-go/question_api"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Home Page")
}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/questions", api.GetQuestions).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	fmt.Println("server running")
	handleRequests()
}
