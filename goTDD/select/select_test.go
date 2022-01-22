package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slow := "http://www.facebook.com"
	fast := "http://www.quii.co.uk"
	want := fast
	got, _ := Racer(slow, fast)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestRacerHttp(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slow := slowServer.URL
	fast := fastServer.URL
	want := fast
	got, _ := Racer(slow, fast)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	slowServer.Close()
	fastServer.Close()
}
