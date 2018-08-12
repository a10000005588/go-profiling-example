package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
)

type counter struct {
	count int
}

var memprofile = flag.String("memprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()

	f, err := os.Create(*memprofile)
	if err != nil {
		panic(err)
	}

	pprof.WriteHeapProfile(f) // start memory profiling

	var wg sync.WaitGroup
	wg.Add(3)

	counter := counter{}

	counter.AddOne()
	counter.AddBillion()
	counter.AddBillion2()

	fmt.Println(counter.count)
}

func (c *counter) AddOne() {
	c.count++
}

func (c *counter) AddBillion() {
	for i := 0; i < 10000; i++ {
		for j := 0; j < 100000; j++ {
			c.count++
		}
	}
}

func (c *counter) AddBillion2() {
	for i := 0; i < 10000; i++ {
		for j := 0; j < 100000; j++ {
			c.count++
		}
	}
}
