package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42")
	}

	return arg + 3, nil
}

// Custom Error Struct
// implement Error interface, return fmt.Sprintf
// return custom error point

type ownError struct {
	arg  int
	prob string
}

func (e *ownError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &ownError{arg, "Can't work with 42"}
	}

	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	_, e := f2(42)
	// TODO：type assertion, to get the data in custom error
	if ae, ok := e.(*ownError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
