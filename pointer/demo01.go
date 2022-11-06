//go:build ignore

package main

import "fmt"

func main() {
	a := 10
	fmt.Println("a value: ", a) //10
	fmt.Printf("a: %T\n", a)    //int
	fmt.Printf("a: %p\n\n", &a) //mem addr

	var p1 *int
	fmt.Println(p1) //nil
	p1 = &a
	fmt.Println("p1 value is: ", p1) //mem addr a
	fmt.Printf("p1: %v\n\n", &p1)    //mem addr p1

	a = 100
	fmt.Println(a)           //100
	fmt.Printf("%p\n\n", &a) //mem addr a

	*p1 = 200
	fmt.Println(a) //200
	var p2 **int
	fmt.Println(p2)                   //null
	fmt.Printf("&p2 is: %p\n\n", &p2) //men addr p2?

	p2 = &p1
	fmt.Printf("&p1 is: %p\n\n", &p1) //men addr p1?
	fmt.Printf("&p2 is: %p\n\n", &p2) //men addr p2?

	fmt.Printf("%T,%T,%T\n", a, p1, p2) //int pointer pointer
	fmt.Println("p2 is: ", p2)          //men addr p1
	fmt.Println("*p2: ", *p2)           //men addr a
	fmt.Println("**p2: ", **p2)         //
}
