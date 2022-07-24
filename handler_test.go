package handler

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := map[string]struct {
		body       string
		expected   response
		statusCode int
	}{
		"happy path": {
			body: "this is four words",
			expected: response{
				WordCount:   4,
				UniqueCount: 4,
				MaxWord:     5,
				AvgWord:     3.75,
				// TODO IP
			},
		},
		"empty": {
			body: "",
			expected: response{
				WordCount:   0,
				UniqueCount: 0,
				MaxWord:     0,
				AvgWord:     0.0,
				// TODO ip
			},
		},
		"one word": {
			body: "asdfasdf",
			expected: response{
				WordCount:   1,
				UniqueCount: 1,
				MaxWord:     8,
				AvgWord:     8.0,
				// TODO ip
			},
		},
		"double space": {
			body: "asd  asdfsll asdfasd     fasdf",
			expected: response{
				WordCount:   4,
				UniqueCount: 4,
				MaxWord:     7,
				AvgWord:     5.5,
				// TODO ip
			},
		},
		"japanese": {
			body: "これは4つの言葉です",
			expected: response{
				WordCount:   1,
				UniqueCount: 1,
				MaxWord:     10,
				AvgWord:     10.0,
				// TODO ip
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			in := strings.NewReader(tt.body)
			req := httptest.NewRequest("POST", "http://example.com", in)
			w := httptest.NewRecorder()
			AllStats(w, req)

			resp := w.Result()
			dec := json.NewDecoder(resp.Body)
			var actual response
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
