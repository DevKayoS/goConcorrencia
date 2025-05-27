package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, name string) {
	ticker := time.NewTicker(250 * time.Millisecond).C

	for {
		select {
		case <-ticker:
			fmt.Println(name, "tick")
		case <-ctx.Done():
			fmt.Println("Recieved signal to finish... exiting...", ctx.Err().Error())
			fmt.Println(name, "Recieved signal to finish... exiting...")
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go worker(ctx, "Worker 1")
	go worker(ctx, "Worker 2")

	time.Sleep(12 * time.Second)
	fmt.Println("Main exiting")
}
