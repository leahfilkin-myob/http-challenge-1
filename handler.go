package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

var stats response
var allWords []string

func TotalWordCount(t []string) int {
	return len(t)
}

func TotalUniqueWordCount(t []string) int {
	counts := make(map[string]int)
	uniqueCount := 0
	for _, v := range t {
		counts[v]++
	}
	log.Printf("Counts: %v", counts)
	for i, _ := range counts {
		if counts[i] < 2 {
			uniqueCount++
		}
	}
	return uniqueCount
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

	allWords = append(allWords, text...)
	stats.WordCount = TotalWordCount(allWords)
	stats.UniqueCount = TotalUniqueWordCount(allWords)
	stats.MaxWord = MaximumWordLength(allWords)
	stats.AvgWord = AverageWordLength(allWords)

	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		panic(err)
	}
}

func GlobalStats(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	if err := enc.Encode(stats); err != nil {
		panic(err)
	}
}

/*func (h *GlobalStats) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.mu.Lock()

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

	h.rt.WordCount += resp.WordCount
	h.rt.UniqueCount += resp.UniqueCount
	h.rt.MaxWord += resp.MaxWord
	h.rt.AvgWord += resp.AvgWord

	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		panic(err)
	}

	h.mu.Unlock()
}
*/
