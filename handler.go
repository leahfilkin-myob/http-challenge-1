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
		if maxCount < len(v) {
			maxCount = len(v)
		}
	}
	return maxCount
}

func AverageWordLength(t []string) float64 {
	totalLetterCount := 0.0
	for _, v := range t {
		totalLetterCount += float64(len(v))
	}
	return totalLetterCount / float64(len(t))
}
func AllStats(w http.ResponseWriter, req *http.Request) {
	/*	t, err := DecodeRequestBody(req.Body, w)
	 */
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

/*func DecodeRequestBody(body io.ReadCloser, w http.ResponseWriter) (text, error) {
	var t text
	dec := json.NewDecoder(body)
	err := dec.Decode(&t)
	if err != nil {
		log.Printf("%v", err)
		fmt.Fprintf(w, "Incorrect text format\n")
	}
	return t, err
}
*/
