//go:build ignore

package main

import "fmt"

func main() {
	fun1()
	fun2 := fun1
	fun2()

	func() {
		fmt.Println("this is a noname")
	}()

	fun3 := func() {
		fmt.Println("this is a noname too")
	}
	fun3()
	fun3()

	func(a, b int) {
		fmt.Println(a, b)
	}(1, 2)

	res1 := func(a, b int) int {
		return a + b
	}(10, 20)
	fmt.Println(res1)

	res2 := func(a, b int) int {
		return a + b
	}
	fmt.Println(res2)

	fmt.Println(res2(100, 200))

}

func fun1() {
	fmt.Println("this is fun1")
}
