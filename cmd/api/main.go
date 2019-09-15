package main

import (
	"fmt"
	"log"
	"net/http"

	app "github.com/luistm/banksaurus/cmd/bscli/application"
)

func version(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, app.Version+"\n")
}

func main() {
	http.HandleFunc("/version", version)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
