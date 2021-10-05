package main

import (
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, my name is Roger Drdn")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:5000", nil)
}
