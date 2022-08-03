package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
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

type Server struct {
	allWords []string
	resp     response
}

func NewServer() *Server {
	return &Server{
		[]string{},
		response{}}
}

func wordCount(t []string) int {
	return len(t)
}

func uniqueCount(t []string) int {
	unique := make(map[string]struct{})
	for _, v := range t {
		if _, ok := unique[v]; !ok {
			unique[v] = struct{}{}
		}
	}
	return len(unique)
}

func maxLength(t []string) int {
	maxCount := 0
	for _, v := range t {
		if maxCount < len([]rune(v)) {
			maxCount = len([]rune(v))
		}
	}
	return maxCount
}

func avgLength(t []string) float64 {
	totalLetterCount := 0.0
	for _, v := range t {
		totalLetterCount += float64(len([]rune(v)))
	}
	if float64(len(t)) == 0.0 {
		return 0
	}
	return math.Round(totalLetterCount/float64(len(t))*100) / 100
}
func (s *Server) AllStats(w http.ResponseWriter, req *http.Request) {
	t, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	defer req.Body.Close()

	text := strings.Fields(string(t))
	resp := response{
		wordCount(text),
		uniqueCount(text),
		maxLength(text),
		avgLength(text),
		req.RemoteAddr,
	}

	s.allWords = append(s.allWords, text...)

	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
}

func (s *Server) GlobalStats(w http.ResponseWriter, req *http.Request) {
	s.resp.WordCount = wordCount(s.allWords)
	s.resp.UniqueCount = uniqueCount(s.allWords)
	s.resp.MaxWord = maxLength(s.allWords)
	s.resp.AvgWord = avgLength(s.allWords)

	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)

	if err := enc.Encode(s.resp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	log.Printf("Resp in method: %v", s.resp)
}
