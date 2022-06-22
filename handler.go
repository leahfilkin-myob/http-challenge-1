package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type text struct {
	TextData string `json:"text"`
}

func TotalWordCount(w http.ResponseWriter, req *http.Request) {
	t, err := DecodeRequestBody(req.Body, w)
	if err == nil {
		count := len(strings.Fields(t.TextData))
		log.Printf("Found word count: %v", count)
		fmt.Fprintf(w, "Word count: %v\n", count)
	}
}

func TotalUniqueWordCount(w http.ResponseWriter, req *http.Request) {
	t, err := DecodeRequestBody(req.Body, w)
	if err == nil {
		words := strings.Fields(t.TextData)
		unique := make(map[string]struct{})
		for _, v := range words {
			if _, ok := unique[v]; !ok {
				unique[v] = struct{}{}
			}
		}
		count := len(unique)
		log.Printf("Found unique word count: %v", count)
		fmt.Fprintf(w, "Unique word count: %v\n", count)
	}
}

func MaximumWordLength(w http.ResponseWriter, req *http.Request) {
	t, err := DecodeRequestBody(req.Body, w)
	if err == nil {
		words := strings.Fields(t.TextData)
		maxCount := 0
		for _, v := range words {
			if maxCount < len(v) {
				maxCount = len(v)
			}
		}
		log.Printf("Found maximum word length: %v", maxCount)
		fmt.Fprintf(w, "Maximum word length: %v\n", maxCount)
	}
}

func AverageWordLength(w http.ResponseWriter, req *http.Request) {
	t, err := DecodeRequestBody(req.Body, w)
	if err == nil {
		words := strings.Fields(t.TextData)
		totalLetterCount := 0.0
		for _, v := range words {
			totalLetterCount += float64(len(v))
		}
		var av = totalLetterCount / float64(len(words))
		log.Printf("Found average word length: %v", av)
		fmt.Fprintf(w, "Average word length: %v\n", av)
	}
}

func DecodeRequestBody(body io.ReadCloser, w http.ResponseWriter) (text, error) {
	var t text
	dec := json.NewDecoder(body)
	err := dec.Decode(&t)
	if err != nil {
		log.Printf("%v", err)
		fmt.Fprintf(w, "Incorrect text format\n")
	}
	return t, err
}
