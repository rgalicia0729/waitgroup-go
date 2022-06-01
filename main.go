package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(2 * time.Second)
	fmt.Printf("Id: %d\n", i)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go doSomething(i, &wg)
	}

	wg.Wait()
}
