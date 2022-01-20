package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	Greet(os.Stdout, "Some")
	http.HandleFunc("/", MyGreetHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, r.URL.Query().Get("name"))
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "hello %s", name)
}
