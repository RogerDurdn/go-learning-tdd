package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye/", goodBye)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func hello(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	fmt.Fprint(res, "Hello, my name is a called:", giveName(name))
}

func goodBye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	fmt.Fprint(res, "Goodbye - ", giveName(name))
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "This is the Home Page")

}
func giveName(name string) string {
	if name == "" {
		name = "some randon name"
	}
	return name
}
