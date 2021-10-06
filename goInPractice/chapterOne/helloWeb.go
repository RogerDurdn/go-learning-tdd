package main

import (
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, my name is Roger Drdn")
}

func main() {
	fmt.Println("Hello, my name is Roger Drdn")
}
