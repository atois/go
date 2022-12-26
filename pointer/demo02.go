//go:build ignore

package main

import "fmt"

func main() {
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Printf("arr1: %v\n", arr1)

	var p1 *[4]int
	p1 = &arr1
	fmt.Println(p1)
	fmt.Printf("p1: %p\n", p1)
	fmt.Printf("&p1: %p\n", &p1)

	(*p1)[0] = 100
	fmt.Printf("arr1: %v\n", arr1)

	p1[0] = 200
	fmt.Printf("arr1: %v\n", arr1)

	a := 1
	b := 2
	c := 3
	d := 4
	arr2 := [4]int{a, b, c, d}
	arr3 := [4]*int{&a, &b, &c, &d}
	fmt.Printf("arr2: %v\n", arr2)
	fmt.Printf("arr3: %v\n", arr3)

	arr2[0] = 100
	fmt.Printf("arr2: %v\n", arr2)
	fmt.Printf("a: %v\n", a)

	*arr3[0] = 200
	fmt.Printf("arr3: %v\n", arr3)
	fmt.Printf("a: %v\n", a)

	b = 1000
	fmt.Printf("arr2: %v\n", arr2)
	fmt.Printf("arr3: %v\n\n", arr3)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(*arr3[i])
	}
}
