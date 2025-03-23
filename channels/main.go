package main

import (
	"fmt"
)

// This function receives from inChan and sends to outChan
func processStrings(inChan <-chan string, outChan chan<- string) {
	for msg := range inChan {
		processed := fmt.Sprintf("Processed: %s", msg)
		outChan <- processed
	}
	close(outChan) // Always a good idea to close when done
}

func main() {
	in := make(chan string)
	out := make(chan string)

	// Start the processor in a goroutine
	go processStrings(in, out)

	// Send some input strings
	go func() {
		in <- "hello"
		in <- "world"
		close(in) // Important to close when done sending
	}()

	// Read processed strings
	for result := range out {
		fmt.Println(result)
	}
}
