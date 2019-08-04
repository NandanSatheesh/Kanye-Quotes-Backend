package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

type QuotesData []string

func main() {

	db, err := sql.Open("mysql", "<username>:<password>@/<database_name>?charset=utf8")
	checkErr(err)

	fmt.Println(db)

	file, err := ioutil.ReadFile("quotes.json")
	checkErr(err)

	var data QuotesData

	_ = json.Unmarshal([]byte(file), &data)

	for i, quote := range data {

		query := "INSERT INTO quotes (data) VALUES (\"" + quote + "\")"
		fmt.Println(query)
		stmt, err := db.Query(query)
		checkErr(err)

		fmt.Println(i, " - > ", stmt)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
