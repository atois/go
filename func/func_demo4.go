//go:build ignore

package main

import "fmt"

func main() {
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Printf("arr1: %v\n", arr1)
	fun1(arr1)
	fmt.Printf("arr1: %v\n", arr1)

	fmt.Println("--------------------------")

	s1 := []int{1, 2, 3, 4}
	fmt.Printf("s1: %v\n", s1)
	fun2(s1)
	fmt.Printf("s1: %v\n", s1)
}
func fun2(s2 []int) {
	fmt.Printf("before %v\n", s2)
	s2[0] = 100
	fmt.Printf("after %v\n", s2)
}

func fun1(arr2 [4]int) {
	fmt.Printf("arr2: %v\n", arr2)
	arr2[0] = 100
	fmt.Printf("arr2: %v\n", arr2)
}
