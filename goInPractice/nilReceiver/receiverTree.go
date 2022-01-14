package main

import "log"

type IntTree struct {
	val int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree  {
	if it == nil{
		return &IntTree{val: val}
	}
	if val < it.val{
		it.left = it.left.Insert(val)
	}else if val > it.val{
		it.left = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int, dir string) (bool, string) {
	switch {
	case it == nil:
		return false, ""
	case val < it.val:
		return  it.left.Contains(val, "left")
	case val > it.val:
		return  it.right.Contains(val, "right")
	default:
		return true, dir
	}
}

func main()  {
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	log.Println(it.Contains(2, ""))
	log.Println(it.Contains(8,""))

}