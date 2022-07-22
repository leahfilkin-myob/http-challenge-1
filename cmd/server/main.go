package main

import (
	handler "github.com/leahfilkin-myob/http-challenge-1"
	"log"
	"net/http"
)

func main() {
	log.Print("Starting handler...")
	http.HandleFunc("/", handler.AllStats)

	http.ListenAndServe(":8090", nil)
}
