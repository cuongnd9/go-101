package main

import (
	"fmt"
)

func main() {
	// Unbuffered channel
	unbuffered := make(chan int)

	// Start a goroutine to send a value to the unbuffered channel
	go func() {
		unbuffered <- 10
		unbuffered <- 20
	}()

	// Receive the value from the unbuffered channel
	for i := 0; i < 2; i++ {
		value := <-unbuffered
		fmt.Println("Unbuffered channel:", value)
	}

	// Buffered channel with a capacity of 1
	buffered := make(chan int, 2)

	// Send a value to the buffered channel
	buffered <- 100
	buffered <- 200

	// Receive the value from the buffered channel
	for i := 0; i < 2; i++ {
		value := <-buffered
		fmt.Println("Buffered channel:", value)
	}
}
