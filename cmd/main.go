package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting \n", id)
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("Worker %d done \n", id)

}

func main() {
	var wg sync.WaitGroup
	const task = 5

	for i := 0; i <= task; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("Main done")

}
