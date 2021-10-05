package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World!", "This flag is for name")

var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "defined to use spanish")
	flag.BoolVar(&spanish, "s", false, "defined to use spanish")
}

func main() {
	flag.Parse()
	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}

	format := "\t-%s: %s (Default: '%s')\n"
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf(format, f.Name, f.Usage, f.DefValue)
	})
}
