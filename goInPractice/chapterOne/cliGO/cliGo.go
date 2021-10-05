package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "Hello World!"
	app.Usage = "Print the Hello World! word"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "name, n", Value: "World", Usage: "Who to say hello world to",
		},
		cli.BoolFlag{
			Name: "spanish, s", Usage: "To see the text on spanish",
		},
	}
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name")
		spanish := c.GlobalBool("spanish")
		if spanish {
			fmt.Printf("Hola %s!\n", name)
		} else {
			fmt.Printf("Hello %s!\n", name)
		}
		return nil
	}
	app.Run(os.Args)
}
