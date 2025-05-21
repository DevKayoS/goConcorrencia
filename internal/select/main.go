package main

import (
	"log"
	"time"
)

func main() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	for i, ch := range chans {
		go func(i int, ch chan<- int) {
			for {
				time.Sleep(time.Duration(i+2) * time.Second) // 2, 3
				ch <- i + 1
			}
		}(i, ch)
	}

	for i := 0; i < 20; i++ {
		select {
		case v := <-chans[0]:
			log.Println("Got on chanel 1: ", v)
		case v := <-chans[1]:
			log.Println("Got on chanel 2: ", v)
		}
	}
}
