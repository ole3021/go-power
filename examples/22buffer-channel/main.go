package main

import (
	"fmt"
)

func main() {
	// buffer channel
	bufferMessage := make(chan string, 2)

	bufferMessage <- "bufferMsg1"
	bufferMessage <- "bufferMsg2"
	// Buffer message exceed the maxium number will cause deadlock.

	fmt.Println(<-bufferMessage)
	fmt.Println(<-bufferMessage)
}
