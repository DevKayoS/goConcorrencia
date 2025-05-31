package main

import "fmt"

func main() {
	ch := make(chan int)

	ch <- 0
	n, ok := <-ch

	fmt.Println("Got a value:", n, "| Is this chanel is closed?", !ok)
	close(ch)

	n, ok = <-ch
	fmt.Println("Got a value:", n, "| Is this chanel is closed?", !ok)
}
