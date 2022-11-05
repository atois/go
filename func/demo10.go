//go:build ignore

package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("a: %T\n", a)

	b := [4]int{1, 2, 3, 4}
	fmt.Printf("b: %T\n", b)

	c := []int{1, 2, 3, 4}
	fmt.Printf("c: %T\n", c)

	d := make(map[int]string)
	fmt.Printf("d: %T\n", d)

	fmt.Printf("fun1: %T\n", fun1)
	fmt.Printf("fun2: %T\n", fun2)
	fmt.Printf("fun3: %T\n", fun3)
}

func fun1() {}

func fun2(a int) int {
	return 0
}

func fun3(a float64, b, c int) (int, int) {
	return 0, 0
}

func fun4(a, b string, c, d int) (string, int, float64) {
	return "", 0, 0
}
