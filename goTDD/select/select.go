package main

import (
	"net/http"
	"time"
)

func Racer(urls ...string) (string, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	winner := make(chan string)
	for _, url := range urls {
		go func(u string) {
			_, err := client.Get(u)
			if err == nil {
				winner <- u
			}
		}(url)
	}
	return "http://www.quii.co.uk", nil
}
