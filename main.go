package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// The data struct for the decoded data
// Notice that all fields must be exportable!
type movies struct {
	ID      string `json:"_id"`
	Movie   string `json:"movie"`
	Salary  int    `json:"salary"`
	Watched bool   `json:"watched"`
}
type books struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
	Pages  int    `json:"pages"`
}

var pathm = "./task/mongoapi-go/mongo.json"
var paths = "./kar/sql.json"

func main() {

	mongo, err := ioutil.ReadFile(pathm)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var m movies
	err = json.Unmarshal(mongo, &m)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	sql, err := ioutil.ReadFile(paths)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var s books
	err = json.Unmarshal(sql, &s)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	log.Printf("id movie: %s\n", m.ID)
	log.Printf("movie: %s\n", m.Movie)
	log.Printf("salary: %d\n\n\n", m.Salary)
	log.Printf("id book: %d\n", s.ID)
	log.Printf("title: %s\n", s.Title)
	log.Printf("pages: %d\n", s.Pages)
	var sum int
	sum = m.Salary + s.Pages
	log.Printf("\nSum:%d\n", sum)
}
