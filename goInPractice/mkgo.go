package main

import (
	"flag"
	"fmt"
	"os"
)

var name = flag.String("n", "", "Flag for name dir and field")

func main() {
	flag.Parse()
	pwd, err := os.Getwd()
	dirPath := pwd + "/" + *name
	err = os.Mkdir(dirPath, 0755)
	fullPath := dirPath + "/" + *name + ".go"
	_, err = os.Create(fullPath)
	err = os.WriteFile(fullPath, getTemplate(), 0644)

	manageError(err)
	fmt.Println("fullPath:", fullPath)
}

func manageError(err error) {
	if err != nil {
		panic(err)
	}
}

func getTemplate() []byte {
	return []byte("package main\n import (\n) \nfunc main() {\n}\n")
}
