package main;

import (
	"fmt"
	"log"
	"net/http"
);

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Request: "+r.URL.Path);
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:]);
}

func main() {
	http.HandleFunc("/",handler);
	log.Print("Listening on port 8080");
	log.Fatal(http.ListenAndServe(":8080",nil));
}
