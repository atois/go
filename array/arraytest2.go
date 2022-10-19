//go:build ignore

package main

import (
	"fmt"
)

func main() {
	arr12 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr12[0])
	fmt.Println(arr12[1])
	fmt.Println(arr12[2])
	fmt.Println(arr12[3])
	fmt.Println(arr12[4])

	fmt.Println("--------------------")

	for i := 0; i < len(arr12); i++ {
		arr12[i] = i*2 + 1
		fmt.Println(arr12[i])
	}
	fmt.Println(arr12)

	fmt.Println("-----------------")

	for index, value := range arr12 {
		fmt.Printf("number is : %d, value is : %d\n", index, value)
	}

}
