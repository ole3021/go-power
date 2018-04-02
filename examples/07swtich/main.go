package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's workday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's morning")
	default:
		fmt.Println("It's afternoon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("It's Bool")
		case int:
			fmt.Println("It's Int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("other")
}
