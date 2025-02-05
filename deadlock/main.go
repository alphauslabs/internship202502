package main

import (
	"fmt"
	"sync"
)

func main() {
	var w sync.WaitGroup
	ch := make(chan int)

	w.Add(2) // Add 2 to the WaitGroup counter

	// run receiver on a separate goroutine
	go receiver(ch, &w)

	// Try to uncomment to fix the deadlock
	// go sender(ch, &w)

	w.Wait() // Wait until the WaitGroup counter is 0
}

func sender(ch chan int, w *sync.WaitGroup) {
	defer func() { // defer will be called when the function is about to return, useful for cleaning up resources
		// It also important to call Done() here to signal that this goroutine is done. This will decrement the WaitGroup counter.
		// Try to comment w.Done() and it will also cause deadlock
		w.Done()
		// Uncomment also this one to close channel
		// This is to signal receiver that sending is done.
		// close(ch)
	}()
	for i := 0; i < 10; i++ {
		ch <- i // Send i to our channel
	}
}

func receiver(ch chan int, w *sync.WaitGroup) {
	defer func() {
		w.Done()
	}()

	for v := range ch { // without closing the channel this will wait infinitely
		fmt.Println("Received ", v)
	}
}
