package main

import (
	"os"
	"projectsBook/goTDD/maths/clockface"
	"time"
)

func main() {

	clockTime := time.Now()
	clockface.SVGWriter(os.Stdout, clockTime)
}
