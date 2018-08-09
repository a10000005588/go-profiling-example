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

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()

	f, err := os.Create(*cpuprofile)
	if err != nil {
		panic(err)
	}

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	var wg sync.WaitGroup
	wg.Add(3)

	counter := counter{}

	var addOne = counter.AddBillion1
	var addMillion = counter.AddBillion2
	var addBillion = counter.AddBillion3

	go func() {
		addOne()
		wg.Done()
	}()

	go func() {
		addMillion()
		wg.Done()
	}()

	go func() {
		addBillion()
		wg.Done()
	}()
	wg.Wait()

	fmt.Println(counter.count)
}

func (c *counter) AddBillion1() {
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

func (c *counter) AddBillion3() {
	for i := 0; i < 10000; i++ {
		for j := 0; j < 100000; j++ {
			c.count++
		}
	}
}
