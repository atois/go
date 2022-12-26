//go:build ignore

package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	fmt.Printf("arr: %v\n", arr)

	fmt.Println("--------------------------")

	s2 := []int{1, 2, 3, 4}
	fmt.Println(s2)
	fmt.Printf("%T,%T\n", arr, s2)

	s3 := make([]int, 3, 8)
	fmt.Printf("s3: %v\n", s3)
	fmt.Printf("capcity: %d\n len: %d\n", cap(s3), len(s3))

	//append
	s4 := make([]int, 0, 5)
	fmt.Println(s4)
	s4 = append(s4, 1, 2)
	fmt.Println(s4)

	s4 = append(s4, s3...)
	fmt.Println(s4)

	for i := 0; i < len(s4); i++ {
		fmt.Println(s4[i])

	}

	for i, v := range s4 {
		fmt.Printf("%d----.%d\n", i, v)
	}

}
