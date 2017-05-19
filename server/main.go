package main

import (
	"github.com/mh-orange/tuner/server"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	router := server.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
