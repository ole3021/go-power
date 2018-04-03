package main

import (
	"fmt"
)

func plus(x int, y int) int {
	return x + y
}

func plusPlus(x, y, z int) int {
	return x + y + z
}

func main() {
	sum1 := plus(3, 4)
	fmt.Println("3+4=", sum1)

	sum2 := plusPlus(1, 2, 4)
	fmt.Println("1+2+4=", sum2)

	a, b := vals()
	fmt.Println("int a:", a)
	fmt.Println("string b:", b)

	_, c := vals()
	fmt.Println("string c:", c)

	sum(1, 2, 3, 4, 5, 6, 7)

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	sum(nums...)
}

// return multiple values
func vals() (int, string) {
	return 3, "test"
}

// variadic arguments
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("sum result", total)
}
