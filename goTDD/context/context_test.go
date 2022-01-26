package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
}

func (s *SpyStore) Fetch(context context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, character := range s.response {
			select {
			case <-context.Done():
				log.Println("spy store got cancelled") // only run if ctx is cancelled initially
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(character)
			}
		}
		data <- result
	}()
	select {
	case <-context.Done():
		log.Println("spy store got cancelled")
		return "", context.Err()
	case result := <-data:
		return result, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (w *SpyResponseWriter) Header() http.Header {
	log.Println("hit the header")
	w.written = true
	return nil
}

func (w *SpyResponseWriter) Write([]byte) (int, error) {
	log.Println("hit the write")
	w.written = true
	return 0, errors.New("not implemented")
}

func (w *SpyResponseWriter) WriteHeader(int) {
	log.Println("hit the writeHeader")
	w.written = true
}

func TestServer(t *testing.T) {
	data := "hello, world!"

	t.Run("happy path", func(t *testing.T) {
		store := &SpyStore{response: data}
		svr := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("Error on server: got %s, want %s", response.Body.String(), data)
		}
	})

	t.Run("cancel Store work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(40*time.Millisecond, cancel)

		request = request.WithContext(cancellingCtx)
		spyResponse := &SpyResponseWriter{}

		svr.ServeHTTP(spyResponse, request)

		if spyResponse.written {
			t.Error("A response should not have been written")
		}
	})
}
