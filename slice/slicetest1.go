//go:build ignore

package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	fmt.Println(s1)

	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1))
	fmt.Printf("s1: %p\n", s1)

	fmt.Println("------------------------")

	s1 = append(s1, 4, 5)
	fmt.Println(s1)
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1))
	fmt.Printf("s1: %p\n", s1)

}
