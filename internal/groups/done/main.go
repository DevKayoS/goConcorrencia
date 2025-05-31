package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("terminou")
		done <- 0
	}()

	<-done
}
