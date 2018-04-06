package main

import (
	"fmt"
)

func main() {
	// create channel with make(chan typeOfMessage)
	message := make(chan string)

	go func() {
		// Send message
		message <- "ping"
	}()

	// receive message
	// receive message is synchronously
	msg := <-message
	fmt.Println(">>> Message from goroutine:", msg)
}
