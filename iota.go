package main

import "fmt"

func main() {
	const (
		a1 = iota
		a2 = iota
		a3 = iota
		_
		a4 = iota
	)

	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
	fmt.Printf("a4: %v\n", a4)

}
