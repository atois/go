//go:build ignore

package main

import "fmt"

func main() {
	res1 := increment()
	fmt.Printf("res1: %T\n", res1)
	fmt.Println(res1)

	v1 := res1()
	fmt.Println(v1)

	v2 := res1()
	fmt.Println(v2)
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())

	res2 := increment()
	fmt.Println(res2())
	v3 := res2()
	fmt.Println(v3)

	fmt.Println("------------------")

	fmt.Println(res2())
	fmt.Println(res1())
}
func increment() func() int {
	i := 0
	fun := func() int {
		i++
		return i
	}
	return fun
}
