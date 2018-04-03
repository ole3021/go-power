package main

import (
	"fmt"
	"math"
)

// interface are named collections of method signatures
// interface implement on struct
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	widht, height float64
}

type circle struct {
	radius float64
}

// interface implementation can not use pointer
func (r rect) area() float64 {
	return r.widht * r.height
}

func (r rect) perim() float64 {
	return 2*r.widht + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{widht: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}