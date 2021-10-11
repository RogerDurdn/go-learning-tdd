package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := rand.Intn(4)
	for a < 10 {
		if a%11 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("do something when the loop completes normally")
done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)
}
