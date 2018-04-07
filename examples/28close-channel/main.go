package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 3)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			// more will be false if jobs has been closed
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("sent job:", i)
	}
	close(jobs) // make channel have no more values
	fmt.Println("sent all jobs")

	<-done
}
