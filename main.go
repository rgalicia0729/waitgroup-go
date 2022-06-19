package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomething(i int, t time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(t)
	fmt.Printf("Se ejecuto gorutina %v\n", i)
}

func main() {
	durations := []time.Duration{
		time.Second * 2,
		time.Second * 2,
		time.Second * 1,
		time.Second * 3,
	}

	var wg sync.WaitGroup

	for i, duration := range durations {
		wg.Add(1)

		func() {
			go doSomething(i+1, duration, &wg)
		}()
	}

	wg.Wait()
}
