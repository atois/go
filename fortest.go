package main

import "fmt"

func main() {
	count := 10
	for i := 0; i < count; i++ {
		fmt.Printf("i=%v\n", i)
		fmt.Printf("i: %v\n", i)
	}

}
