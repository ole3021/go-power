package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"bob", 30})
	fmt.Println(person{name: "Alice", age: 19})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "John", age: 29})

	s := person{name: "Sean", age: 23}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 44
	// dot the structure, will automatically dereference the pointer
	fmt.Println(sp.age)
	// struct are mutable
	fmt.Println(s.age)
}
