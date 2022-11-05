//go:build ignore

package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", oper)

}

func add(a, b int) int {
	return a + b
}

func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun)
	return 0
}
