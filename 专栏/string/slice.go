package main

import "fmt"

func main() {
	s := "hello world"
	a := 2
	b := 5

	fmt.Printf("s[a]: %c\n", s[a])
	fmt.Printf("s[a:b]: %v\n", s[a:b])
	fmt.Printf("s[a:]: %v\n", s[a:])
	fmt.Printf("s[:b]: %v\n", s[:b])
}
