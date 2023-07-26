package question_api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Question struct {
	Category   string   `json:"category"`
	Id         string   `json:"id"`
	Tags       []string `json:"tags"`
	Difficulty string   `json:"difficulty"`
	Regions    []string `json:"regions"`
	IsNiche    bool     `json:"isNiche"`
	Question   struct {
		Text string `json:"text"`
	} `json:"question"`
	CorrectAnswer    string   `json:"correctAnswer"`
	IncorrectAnswers []string `json:"incorrectAnswers"`
	Type             string   `json:"type"`
}

type Questions []Question

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	baseUrl := "https://the-trivia-api.com"
	resource := "/v2/questions"

	// Add Parameters
	parameters := url.Values{}
	parameters.Add("categories", "history")

	u, _ := url.ParseRequestURI(baseUrl)
	u.Path = resource
	u.RawQuery = parameters.Encode()
	urlStr := fmt.Sprintf("%v", u)

	fmt.Println(urlStr)
	res, err := http.Get(urlStr)
	if err != nil {
		log.Fatal(fmt.Errorf("hitting endpoint %s: %v", urlStr, err))
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("panic!!")
	}

	var questions Questions
	json.Unmarshal(body, &questions)

	// Use json.Decode for reading streams of JSON data
	if err := json.NewEncoder(w).Encode(questions); err != nil {
		log.Println(err)
	}
}
