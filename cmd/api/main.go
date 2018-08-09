package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/luistm/banksaurus/app"
)

func version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, app.Version+"\n")
}

func main() {
	http.HandleFunc("/version", version)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
