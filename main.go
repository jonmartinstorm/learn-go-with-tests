package main

import (
	"log"
	"net/http"

	"github.com/jonmartinstorm/learn-go-with-tests/dependencyinjection"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	dependencyinjection.Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
