package main

import (
	handler "github.com/leahfilkin-myob/http-challenge-1"
	"log"
	"net/http"
)

func main() {
	log.Print("Starting handler...")
	http.HandleFunc("/totalWordCount", handler.TotalWordCount)
	http.HandleFunc("/totalUniqueWords", handler.TotalUniqueWordCount)
	http.HandleFunc("/maximumWordLength", handler.MaximumWordLength)
	http.HandleFunc("/averageWordLength", handler.AverageWordLength)

	http.ListenAndServe(":8090", nil)
}
