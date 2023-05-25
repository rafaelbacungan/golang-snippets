package concurrency

import "fmt"

func ConcurrencySample() {
	go func() {
		fmt.Println("hello world. this is a goroutine")
	}()
	fmt.Println("hello world")

	// channels
	c := make(chan string)
	go func() {
		c <- "hello world. this is a channel"
	}()

	fmt.Println(<-c)
}
