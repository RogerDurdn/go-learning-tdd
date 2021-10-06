package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

var name = flag.String("n", "", "Flag for name dir and fiel")

func main() {
	flag.Parse()
	ex, err := os.Getwd()
	errorM(err)
	exPath := path.Dir(ex)
	dirPath := exPath + "/" + *name
	err = os.Mkdir(dirPath, 0755)
	errorM(err)
	fullPath := dirPath + "/" + *name + ".go"
	fmt.Println("fullPath:", fullPath)
	someFile, err := os.Create(fullPath)
	errorM(err)
	log.Println(someFile)
	someFile.Close()
	err = os.WriteFile(fullPath, getTemplate(), 0644)
	errorM(err)
}

func errorM(err error) {
	if err != nil {
		panic(err)
	}

}
func getTemplate() []byte {
	return []byte("package main\n import (\n) \nfunc main() {\n}\n")
}
