package main

import (
	"fmt"

	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {
	config, err := yaml.ReadFile("yamlConf.yml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Get("path"))
	fmt.Println(config.GetBool("enabled"))

}
