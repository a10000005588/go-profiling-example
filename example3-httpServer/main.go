package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func Counter(wg *sync.WaitGroup) {
	time.Sleep(time.Second)

	var counter int
	for i := 0; i < 1000000; i++ {
		time.Sleep(time.Millisecond * 100)
		counter++
	}
	wg.Done()
}

func main() {
	flag.Parse()

	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Counter(&wg)
	}
	wg.Wait()

	time.Sleep(60 * time.Second)
}
