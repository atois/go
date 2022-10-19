//go:build ignore

package main

import "fmt"

func main() {
	a := 100
	b := 20
	fmt.Printf("(a + b): %v\n", (a + b))
	fmt.Printf("(a - b): %v\n", (a - b))
	fmt.Printf("(a / b): %v\n", (a / b))
	fmt.Printf("(a * b): %v\n", (a * b))

	x := a % b
	fmt.Printf("x: %v\n", x)
}
