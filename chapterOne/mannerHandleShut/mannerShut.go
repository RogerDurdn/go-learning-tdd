package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/braintree/manners"
)

func newHandler() *handler {
	return &handler{}
}

type handler struct{}

func main() {
	handler := newHandler()
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go listenForShutDown(ch)

	fmt.Println("The homepage port is:", os.Getenv("PORT"))
	manners.ListenAndServe("localhost:"+os.Getenv("PORT"), handler)

}

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	shut := query.Get("shut")
	go someRuntine(3000)
	if name == "" {
		name = "Inigo Montoyasssss"
	} else {
		fmt.Fprint(res, "Hello my name is ", name)
	}
	if shut != "" {
		os.Exit(0)
	}
}
func listenForShutDown(ch <-chan os.Signal) {
	<-ch
	manners.Close()
}

func someRuntine(elements int) {
	for i := 0; i < elements; i++ {
		fmt.Print(" ", i, " ")
	}
}
