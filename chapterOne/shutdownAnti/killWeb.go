package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/shutDown", shutDown)
	http.HandleFunc("/", homePage)
	fmt.Println("The homepage port is:", os.Getenv("PORT"))
	http.ListenAndServe("localhost:"+os.Getenv("PORT"), nil)

}
func shutDown(res http.ResponseWriter, req *http.Request) {
	fmt.Println("We are closing everithing")
	os.Exit(0)
}
func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
	}
	fmt.Fprint(res, "The homepage motha fucka")
}
