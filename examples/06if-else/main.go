package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is event")
	} else {
		fmt.Println("7 is not event")
	}

	if 8%4 == 0 {
		fmt.Println("8 is event")
	} else {
		fmt.Println("8 is not event")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
