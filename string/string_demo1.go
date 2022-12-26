//go:build ignore

package main

import (
	"fmt"
)

func main() {
	s1 := "hello"
	s2 := "hello world"
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Printf("s1: %v\n", len(s1))
	fmt.Printf("s2: %v\n", len(s2))

	fmt.Println(s2[0])
	a := 'h'
	b := 104
	fmt.Printf("%c,%c,%c\n", s2[0], a, b)

	for i := 0; i < len(s2); i++ {
		fmt.Println(s2[i])
	}

	fmt.Println("---------------------------")
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%c\t", s1[i])
	}

	//for range
	for _, v := range s2 {
		fmt.Printf("%c\n", v)
	}

	slice1 := []byte{65, 66, 67, 68, 69}
	s3 := string(slice1)
	fmt.Println(s3)

}
