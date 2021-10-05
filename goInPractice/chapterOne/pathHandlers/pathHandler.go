package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	pr := newPathResolver()
	pr.Add("GET /hello", hello)
	pr.Add("* /goodbye/*", goodBye)
	http.ListenAndServe(":"+os.Getenv("PORT"), pr)
}

func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path
	for pattern, handleFunc := range p.handlers {
		if ok, err := path.Match(pattern, check); ok && err == nil {
			handleFunc(res, req)
			return
		} else if err != nil {
			fmt.Fprint(res, err)
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
