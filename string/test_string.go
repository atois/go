//go:build ignore

package main

import "fmt"

func main() {
	var s string = "hello world"
	var s1 = "hello world"
	s3 := "hello world"

	fmt.Printf("s: %v\n", s)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s3: %v\n", s3)

	s4 := `
	line 1
	line 2
	line 3
	`
	fmt.Printf("s4: %v\n", s4)
}
