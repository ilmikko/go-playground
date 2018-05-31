package main

import (
	"net/http"
	//"io/ioutil"
	"fmt"
	//"time"
	//"encoding/json"
	"log"
)

type Task struct {
	Title       string `json:"title"`
	Completed   bool   `json:"completed"`
	DateCreated int64  // time.Now() will be substituted
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos"
	// No blank identifier => use all variables inevitably
	result, error := http.Get(url)
	if error != nil {
		log.Fatal(error) // the log exits the program immedialty unlike fmt.Print..
	}
	// blank identifier => no error catching, and crach risk
	fmt.Println(result.Cookies())
}
