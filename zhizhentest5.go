package main

import "fmt"

func main() {
	p1 := 1
	p2 := &p1
	*p2++
	fmt.Println(p1)
	fmt.Println(*p2)
}
