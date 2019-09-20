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
	const port = ":8080"
	fmt.Println("Starting banksaurus web server on port", port)
	http.HandleFunc("/version", version)
	log.Fatal(http.ListenAndServe(port, nil))
}
