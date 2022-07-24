package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type response struct {
	WordCount   int
	UniqueCount int
	MaxWord     int
	AvgWord     float64
	SourceIP    string
}

func TotalWordCount(t []string) int {
	return len(t)
}

func TotalUniqueWordCount(t []string) int {
	unique := make(map[string]struct{})
	for _, v := range t {
		if _, ok := unique[v]; !ok {
			unique[v] = struct{}{}
		}
	}
	return len(unique)
}

func MaximumWordLength(t []string) int {
	maxCount := 0
	for _, v := range t {
		if maxCount < len([]rune(v)) {
			maxCount = len([]rune(v))
		}
	}
	return maxCount
}

func AverageWordLength(t []string) float64 {
	totalLetterCount := 0.0
	for _, v := range t {
		totalLetterCount += float64(len([]rune(v)))
	}
	if float64(len(t)) == 0.0 {
		return 0
	}
	return totalLetterCount / float64(len(t))
}
func AllStats(w http.ResponseWriter, req *http.Request) {
	t, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	text := strings.Fields(string(t))
	resp := response{
		TotalWordCount(text),
		TotalUniqueWordCount(text),
		MaximumWordLength(text),
		AverageWordLength(text),
		req.RemoteAddr,
	}
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		panic(err)
	}
}
