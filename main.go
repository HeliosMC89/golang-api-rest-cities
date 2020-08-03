package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	dbConnect()
	log.Fatal(http.ListenAndServe(":8000", router))
}
