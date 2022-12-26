//go:build ignore

package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", oper)

	res1 := add(1, 2)
	fmt.Println(res1)

	res2 := oper(10, 20, add)
	fmt.Println(res2)

	res3 := oper(5, 2, sub)
	fmt.Println(res3)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun)
	res := fun(a, b)
	return res
}
