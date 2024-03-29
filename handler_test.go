package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := map[string]struct {
		body       string
		expected   Response
		statusCode int
	}{
		"happy path": {
			body: "this is four words",
			expected: Response{
				WordCount:   4,
				UniqueCount: 4,
				MaxWord:     5,
				AvgWord:     3.75,
				// TODO IP
			},
			statusCode: 200,
		},
		"empty": {
			body: "",
			expected: Response{
				WordCount:   0,
				UniqueCount: 0,
				MaxWord:     0,
				AvgWord:     0.0,
				// TODO ip
			},
			statusCode: 200,
		},
		"one word": {
			body: "asdfasdf",
			expected: Response{
				WordCount:   1,
				UniqueCount: 1,
				MaxWord:     8,
				AvgWord:     8.0,
				// TODO ip
			},
			statusCode: 200,
		},
		"double space": {
			body: "asd  asdfsll asdfasd     fasdf",
			expected: Response{
				WordCount:   4,
				UniqueCount: 4,
				MaxWord:     7,
				AvgWord:     5.5,
				// TODO ip
			},
			statusCode: 200,
		},
		"japanese": {
			body: "これは4つの言葉です",
			expected: Response{
				WordCount:   1,
				UniqueCount: 1,
				MaxWord:     10,
				AvgWord:     10.0,
				// TODO ip
			},
			statusCode: 200,
		},
		"unique words": {
			body: "brown brown brown fox",
			expected: Response{
				WordCount:   4,
				UniqueCount: 2,
				MaxWord:     5,
				AvgWord:     4.5,
				// TODO ip
			},
			statusCode: 200,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			in := strings.NewReader(tt.body)
			req := httptest.NewRequest("POST", "http://example.com", in)
			w := httptest.NewRecorder()
			s := NewServer()
			s.AllStats(w, req)

			resp := w.Result()
			dec := json.NewDecoder(resp.Body)
			var actual Response
			if err := dec.Decode(&actual); err != nil {
				t.Fatalf("failed to decode: %s", err)
			}
			// TODO ask about defers in this context
			defer resp.Body.Close()

			if actual.WordCount != tt.expected.WordCount {
				t.Errorf("got %d, want %d", actual.WordCount, tt.expected.WordCount)
			}
			if actual.UniqueCount != tt.expected.UniqueCount {
				t.Errorf("got %d, want %d", actual.UniqueCount, tt.expected.UniqueCount)
			}
			if actual.MaxWord != tt.expected.MaxWord {
				t.Errorf("got %d, want %d", actual.MaxWord, tt.expected.MaxWord)
			}
			if actual.AvgWord != tt.expected.AvgWord {
				t.Errorf("got %v, want %v", actual.AvgWord, tt.expected.AvgWord)
			}
		})
	}
}

func TestGlobal(t *testing.T) {
	tests := map[string]struct {
		bodies     []string
		expected   Response
		statusCode int
	}{
		"happy path": {
			bodies: []string{
				"this is four words",
				"this is now five words",
			},
			expected: Response{
				WordCount:   9,
				UniqueCount: 6,
				MaxWord:     5,
				AvgWord:     3.67,
				// TODO IP
			},
			statusCode: 200,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			s := NewServer()
			for _, v := range tt.bodies {
				in := strings.NewReader(v)
				req := httptest.NewRequest(http.MethodPost, "/", in)
				s.AllStats(w, req)
			}
			req := httptest.NewRequest(http.MethodGet, "/global", nil)
			x := httptest.NewRecorder()
			s.GlobalStats(x, req)
			resp := x.Result()
			dec := json.NewDecoder(resp.Body)
			var actual Response
			if err := dec.Decode(&actual); err != nil {
				t.Fatalf("failed to decode: %s", err)
			}
			// TODO ask about defers in this context
			defer resp.Body.Close()

			if actual.WordCount != tt.expected.WordCount {
				t.Errorf("got %d, want %d", actual.WordCount, tt.expected.WordCount)
			}
			if actual.UniqueCount != tt.expected.UniqueCount {
				t.Errorf("got %d, want %d", actual.UniqueCount, tt.expected.UniqueCount)
			}
			if actual.MaxWord != tt.expected.MaxWord {
				t.Errorf("got %d, want %d", actual.MaxWord, tt.expected.MaxWord)
			}
			if actual.AvgWord != tt.expected.AvgWord {
				t.Errorf("got %v, want %v", actual.AvgWord, tt.expected.AvgWord)
			}
		})
	}
}

func TestGlobalConcurrent(t *testing.T) {
	tests := map[string]struct {
		body       string
		expected   Response
		statusCode int
	}{
		"happy path": {
			body: "this is four words",
			expected: Response{
				WordCount:   4000,
				UniqueCount: 4,
				MaxWord:     5,
				AvgWord:     3.75,
				// TODO IP
			},
			statusCode: 200,
		},
	}

	for name, tt := range tests {
		s := NewServer()
		numberOfTests := 1000
		t.Run(name, func(t *testing.T) {
			wg := sync.WaitGroup{}
			w := httptest.NewRecorder()
			for i := 0; i < numberOfTests; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					in := strings.NewReader(tt.body)
					req := httptest.NewRequest(http.MethodPost, "/", in)
					s.AllStats(w, req)
				}()
			}
			wg.Wait()
			req := httptest.NewRequest(http.MethodGet, "/global", nil)
			x := httptest.NewRecorder()
			s.GlobalStats(x, req)
			resp := x.Result()
			dec := json.NewDecoder(resp.Body)
			var actual Response
			if err := dec.Decode(&actual); err != nil {
				t.Fatalf("failed to decode: %s", err)
			}
			// TODO ask about defers in this context
			defer resp.Body.Close()

			if actual.WordCount != tt.expected.WordCount {
				t.Errorf("got %d, want %d", actual.WordCount, tt.expected.WordCount)
			}
			if actual.UniqueCount != tt.expected.UniqueCount {
				t.Errorf("got %d, want %d", actual.UniqueCount, tt.expected.UniqueCount)
			}
			if actual.MaxWord != tt.expected.MaxWord {
				t.Errorf("got %d, want %d", actual.MaxWord, tt.expected.MaxWord)
			}
			if actual.AvgWord != tt.expected.AvgWord {
				t.Errorf("got %v, want %v", actual.AvgWord, tt.expected.AvgWord)
			}
		})
	}
}
