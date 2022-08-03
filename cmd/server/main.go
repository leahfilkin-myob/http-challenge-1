package main

import (
	handler "github.com/leahfilkin-myob/http-challenge-1"
	"log"
	"net/http"
)

/*log.Print("Starting handler...")
ch := new(handler.CountHandler)
http.HandleFunc("/hello", handler.Hello)
http.HandleFunc("/health", handler.Health)
http.HandleFunc("/headers", handler.Headers)
http.HandleFunc("/metadata", handler.Metadata)
http.Handle("/count", ch)*/

func main() {
	h := handler.NewServer()
	log.Print("Starting handler...")
	http.HandleFunc("/", h.AllStats)
	http.HandleFunc("/global", h.GlobalStats)

	http.ListenAndServe(":8090", nil)
}
