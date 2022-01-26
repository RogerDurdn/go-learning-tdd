package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {

}

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			log.Println("Error on fetch:", err)
			return // todo: log error however
		}
		fmt.Fprintf(w, data)
	}
}
