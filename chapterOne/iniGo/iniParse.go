package main

import (
	"fmt"

	"gopkg.in/gcfg.v1"
)

func main() {
	config := struct {
		Section struct {
			Enabled bool
			Path    string
		}
	}{}
	err := gcfg.ReadFileInto(&config, "conf.ini")
	if err != nil {
		fmt.Printf("failed to parse the file:%s \n", err)
	}
	fmt.Println(config.Section.Enabled)
	fmt.Println(config.Section.Path)
}
