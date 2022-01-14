package main

import (
 "fmt"
 pnt "projectsBook/goInPractice/packages/formatter"
 "projectsBook/goInPractice/packages/math"
)

func main() {
 num := math.Double(10)
 output := pnt.Format(num)
 fmt.Println(output)
}
