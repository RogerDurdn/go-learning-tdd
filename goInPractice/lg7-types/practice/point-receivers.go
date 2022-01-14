package main

import (
	"fmt"
	"log"
	"time"
)

type Counter struct {
	total int
	lastUpdate time.Time
}

func (c *Counter) increment()  {
	c.total ++
	c.lastUpdate = time.Now()
}

func (c Counter) toString()  {
	log.Println(fmt.Sprintf("f: todal: %d, lastUpdat: %v", c.total, c.lastUpdate.Nanosecond()))
}

func doUpdateWrong(c Counter)  {
	c.increment()
	c.toString()
}

func doUpdateRight(c *Counter)  {
	c.increment()
	c.toString()
}

func main()  {
	count := Counter{1, time.Now()}

	count.increment()
	count.increment()
	count.toString()
	count.increment()
	doUpdateWrong(count)
	count.increment()
	count.toString()
	doUpdateRight(&count)
}