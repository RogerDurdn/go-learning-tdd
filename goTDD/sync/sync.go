package main

import "sync"

func main() {
}

type Counter struct {
	mu    sync.Mutex
	count int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Value() int {
	return c.count
}
