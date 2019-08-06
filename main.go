package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type KanyeQuotes struct {
	QuoteString string `json:"quote"`
}

type QuotesData []string

var dataFromFile []string

func init() {

	dataFromFile = make([]string, 60)

	file, err := ioutil.ReadFile("quotes.json")
	checkErr(err)

	var data QuotesData

	err = json.Unmarshal([]byte(file), &data)
	checkErr(err)

	for i, quote := range data {
		dataFromFile[i] = quote
	}
}

func getDataFromFile() string {

	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(60)

	return dataFromFile[id]
}

func GetRandomQuote(w http.ResponseWriter, r *http.Request) {

	quotesData := KanyeQuotes{
		QuoteString: getDataFromFile(),
	}

	data, err := json.Marshal(quotesData)
	checkErr(err)

	fmt.Println(string(data))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func handleRequest() {
	http.HandleFunc("/quote", GetRandomQuote)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
