package main

import "fmt"

func main() {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	present(v, ok)

	v, ok = m["world"]
	present(v, ok)

	v, ok = m["goodbye"]
	present(v, ok)
}

func present(v int, ok bool) {
	fmt.Printf("value:%d %v\n", v, ok)
}
