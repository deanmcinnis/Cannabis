package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	log.Println("Listening on port localhost:3000..")
	http.ListenAndServe(":3000", nil)
}
