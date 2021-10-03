package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("The homepage port is:", os.Getenv("PORT"))
	http.HandleFunc("/", homePage)
	http.ListenAndServe("localhost:"+os.Getenv("PORT"), nil)

}
func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
	}
	fmt.Fprint(res, "The homepage motha fucka")

}
