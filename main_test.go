package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkEncoder(b *testing.B) {
	w := httptest.NewRecorder()
	r := new(http.Request)
	for n := 0; n < b.N; n++ {
		healthcheckHandlerEncoder(w, r)
	}
}

func BenchmarkMarshal(b *testing.B) {
	w := httptest.NewRecorder()
	r := new(http.Request)
	for n := 0; n < b.N; n++ {
		healthcheckHandlerMarshal(w, r)
	}
}

// To run benchmark tests
// go test -run=^$ -bench=. -benchmem -count=3 -benchtime=5s
