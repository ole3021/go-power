package main

import (
	"fmt"
	"math"
)

const s = "constant value"

func main() {
	fmt.Println(s)

	const n = 40000

	const d = 3e20 / n
	fmt.Println(d, int64(d))

	fmt.Println(math.Sin(n))
}
