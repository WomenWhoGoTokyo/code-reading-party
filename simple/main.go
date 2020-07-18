package main

import (
	"log"
	"net/http"
)

func main() {
	// Set Handlers
	http.HandleFunc("/", HelloHandler)

	// Start
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Bear!"))
}
