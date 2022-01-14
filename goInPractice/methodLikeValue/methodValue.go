package main

import (
	"log"
)

type Adder struct {
	start int
}

func (a Adder) Add(number int) int {
	return a.start + number
}


func main()  {
	myAddr := Adder{10}
	log.Println(myAddr.Add(5))

	myAddFnc := myAddr.Add
	log.Println(myAddFnc(10))

	myAddFnc2 := Adder.Add
	log.Println(myAddFnc2(myAddr, 20))

}