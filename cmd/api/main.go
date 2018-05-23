package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
