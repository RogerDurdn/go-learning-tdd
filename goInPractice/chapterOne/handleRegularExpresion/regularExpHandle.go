package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	rr := newPathResolver()
	rr.Add("GET /hello", hello)
	rr.Add("(GET|HEAD) /goodbye(/?[A-Za-z0-9]*)?", goodBye)
	http.ListenAndServe(":"+os.Getenv("PORT"), rr)
}

func newPathResolver() *regexResolver {

	return &regexResolver{
		handlers: make(map[string]http.HandlerFunc),
		cache:    make(map[string]*regexp.Regexp),
	}
}

type regexResolver struct {
	handlers map[string]http.HandlerFunc
	cache    map[string]*regexp.Regexp
}

func (r *regexResolver) Add(regex string, handler http.HandlerFunc) {
	r.handlers[regex] = handler
	cache, _ := regexp.Compile(regex)
	r.cache[regex] = cache
}

func (r *regexResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path
	for pattern, handleFunc := range r.handlers {
		if r.cache[pattern].MatchString(check) == true {
			handleFunc(res, req)
			return
		}
	}
	http.NotFound(res, req)
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

func giveName(name string) string {
	if name == "" {
		name = "some randon name"
	}
	return name
}
