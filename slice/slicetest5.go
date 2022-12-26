//go:build ignore

package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 0) //len:0 cap:0
	for i := 0; i < len(s1); i++ {
		s2 = append(s2, s1[i])
	}
	fmt.Println(s1)
	fmt.Println(s2)

	s1[0] = 100
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)

	fmt.Printf("------------------copy test------------------------\n")

	s3 := []int{7, 8, 9}
	fmt.Println(s2)
	fmt.Println(s3)

	copy(s3, s2) //s2 --->s3
	fmt.Println(s2)
	fmt.Println(s3)

	copy(s3[1:], s2[2:])
	fmt.Println(s2)
	fmt.Println(s3)

}
