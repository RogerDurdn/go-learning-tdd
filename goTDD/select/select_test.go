package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacerHttp(t *testing.T) {
	t.Run("Select fast connection", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slow := slowServer.URL
		fast := fastServer.URL
		want := fast

		got, err := Racer(slow, fast)
		if err != nil {
			t.Fatal("Not expected error: ", err)
		}
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("Error after seconds of timeout", func(t *testing.T) {
		delayedServer := makeDelayedServer(20 * time.Millisecond)
		defer delayedServer.Close()

		_, err := ConfigurableRacer(delayedServer.URL, delayedServer.URL, 10*time.Millisecond)
		if err == nil {
			t.Errorf("expected error of timeout after seconds of timeout")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
