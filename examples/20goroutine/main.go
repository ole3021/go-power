package main

import (
	"fmt"
)

func f(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Println("last synchronously print2")
	// Press any key to exit
	fmt.Scanln()
	fmt.Println("never execute")
}
