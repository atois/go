package main

import (
	"fmt"
	"reflect"
)

func main() {
	//get address
	var x = 10
	fmt.Printf("before: %p\n", &x)
	//x.p
	x = 100
	fmt.Printf("after: %p\n", &x)

	//address assign to pointer
	var p *int
	p = &x
	fmt.Println(p)

	//get value from pointer
	fmt.Println(*p, reflect.TypeOf(*p))

	*p = 90
	fmt.Println(x)

}
