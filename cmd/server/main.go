package main

import (
	"fmt"
	handler "github.com/leahfilkin-myob/http-challenge-1"
	"log"
	"net/http"
	"time"
)

func main() {

	h := handler.NewServer()
	log.Print("Starting handler...")
	http.HandleFunc("/", h.AllStats)
	http.HandleFunc("/global", h.GlobalStats)

	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			fmt.Println(h.CalculateGlobalStats())
		}
	}()

	http.ListenAndServe(":8090", nil)
}
