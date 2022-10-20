//go:build ignore

package main

import "fmt"

func main() {
	//value
	a1 := [4]int{1, 2, 3, 4}
	a2 := a1
	fmt.Println(a1, a2)
	a1[0] = 100
	fmt.Println(a1, a2)

	//slice
	s1 := []int{1, 2, 3, 4}
	s2 := s1
	fmt.Println(s1, s2)
	s1[0] = 100
	fmt.Println(s1, s2)
}
