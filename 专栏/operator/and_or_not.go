package main

import "fmt"

func main() {
	a := true
	b := false

	r := a && b
	fmt.Printf("r: %v\n", r)

	r = a || b
	fmt.Printf("r: %v\n", r)

	r = !a
	fmt.Printf("r: %v\n", r)
}
