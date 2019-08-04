package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type KanyeQuotes struct {
	QuoteString string `json:"quote"`
}

func GetRandomQuote(w http.ResponseWriter, r *http.Request) {

	quotesData := KanyeQuotes{
		QuoteString: getDataFromDataBase(),
	}

	data, err := json.Marshal(quotesData)
	checkErr(err)

	fmt.Println(string(data))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func getDataFromDataBase() string {
	db, err := sql.Open("mysql", "<username>:<password>@/<database_name>?charset=utf8")
	checkErr(err)

	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(60) + 1

	stmt, err := db.Prepare("SELECT data FROM quotes WHERE id = ?")
	checkErr(err)

	var data string
	defer stmt.Close()

	rows, err := stmt.Query(id)
	checkErr(err)

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&data)
		return data
	}

	return "Something went wrong"

}

func handleRequest() {
	http.HandleFunc("/quote", GetRandomQuote)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	handleRequest()
}
