//go:build ignore

package main

import "fmt"

func main() {
	a := 10
	a += 5
	fmt.Printf("a: %v\n", a)

	b := [4]int{1, 2, 3, 4}
	b[0] = 100
	for i := 0; i < len(b); i++ {
		fmt.Printf("b[i]: %v\n", b[i])
	}
	fmt.Println()

	fmt.Printf("fun1: %T\n", fun1)
	fmt.Println(fun1)

	var c func(int, int)
	fmt.Println(c)
}

func fun1(a, b int) {
	fmt.Printf("a:%d,b:%d\n", a, b)
}

