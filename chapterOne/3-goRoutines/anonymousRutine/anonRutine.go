package main

import (
	"fmt"
	// "runtime"
)

func main() {
	fmt.Println("Outside a goroutine")
	go func() {
		fmt.Println("Inside a goroutine")
	}()
	fmt.Println("Outsite again")
	// runtime.Gosched()
}
