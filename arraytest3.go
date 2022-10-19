//go:build ignore

package main

import (
	"fmt"
)

func main() {
	num := 10
	fmt.Printf("format: %T\n", num)

	arr113 := [4]int{1, 2, 3, 4}
	arr213 := [3]float64{2.15, 3.18, 6.19}
	arr313 := [4]int{5, 6, 7, 8}
	arr413 := [2]string{"hello", "world"}

	fmt.Printf("arr113: %T\n", arr113)
	fmt.Printf("arr213: %T\n", arr213)
	fmt.Printf("arr313: %T\n", arr313)
	fmt.Printf("arr413: %T\n", arr413)

}
