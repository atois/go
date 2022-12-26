package main

import "fmt"

func main() {
	var p *int
	p = new(int)
	fmt.Println(*p)
	*p = 10
	fmt.Println(*p)
}
