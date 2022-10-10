package main

import "fmt"

func main() {
	a := 10
	b := 5

	r := a == b
	fmt.Printf("r: %v\n", r)

	r = a > b
	fmt.Printf("r: %v\n", r)

	r = a < b
	fmt.Printf("r: %v\n", r)

	r = a >= b
	fmt.Printf("r: %v\n", r)

	r = a <= b
	fmt.Printf("r: %v\n", r)

	r = a != b
	fmt.Printf("r: %v\n", r)

}
