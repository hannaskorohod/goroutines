package main

import (
	"fmt"
	"sync"
	"time"
)

func myFunc(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go myFunc(i, &wg)
	}

	wg.Wait()
}
