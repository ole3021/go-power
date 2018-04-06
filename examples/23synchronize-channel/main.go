package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("start working...")
	time.Sleep(time.Second * 3)
	fmt.Println("done")
	done <- true
}

func main() {
	sdone := make(chan bool, 1)

	go worker(sdone)

	<-sdone
}
