package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	done := make(chan bool)
	// timer will provide a channle to notify

	<-timer1.C
	fmt.Println("Timer1 expired")

	timer2 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 expired")
		done <- true
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer2 Stopped")
	}

	select {
	case <-done:
		fmt.Println("print complete")
	// timeout select
	case <-time.After(time.Second * 3):
		fmt.Println("timeouted at 3")
	}
}
