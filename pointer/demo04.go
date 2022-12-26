//go:build ignore

package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("a: %v\n", a)
	fun1(a)
	fmt.Printf("a(fun1): %v\n", a)

	fun2(&a)
	fmt.Printf("a(fun2): %v\n", a)

	arr1 := [4]int{1, 2, 3, 4}
	fmt.Printf("arr1: %v\n", arr1)
	fun3(arr1)
	fmt.Printf("arr1(fun3): %v\n", arr1)

	fun4(&arr1)
	fmt.Printf("arr1(fun4): %v\n", arr1)
}

func fun4(p2 *[4]int) {
	fmt.Printf("p2: %v\n", p2)
	p2[0] = 200
	fmt.Printf("p2: %v\n", p2)
}

func fun3(arr2 [4]int) {
	fmt.Printf("arr2: %v\n", arr2)
	arr2[0] = 100
	fmt.Printf("after arr2: %v\n", arr2)
}
func fun1(num int) {
	fmt.Printf("num: %v\n", num)
	num = 100
	fmt.Printf("num: %v\n", num)
}
func fun2(p1 *int) {
	fmt.Printf("p1: %v\n", *p1)
	*p1 = 200
	fmt.Printf("after p1: %v\n", *p1)
}
