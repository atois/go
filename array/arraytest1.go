//go:build ignore

package main

import "fmt"

func main() {
	var num1 int
	num1 = 100

	num1 = 200
	fmt.Println(num1)

	var arr1 [4]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	fmt.Printf("arr1[3]: %v\n", arr1[3])
}
