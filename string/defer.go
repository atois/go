//go:build ignore

package main

import "fmt"

func main() {
	defer fun1("hello")
	fmt.Println("123456")
	defer fun1("world")
	fmt.Println("king")
}

func fun1(s string) {
	fmt.Println(s)
}
