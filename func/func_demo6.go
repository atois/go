//go:build ignore

package main

import (
	"fmt"
)

func main() {
	a, b, c := fun1()
	fmt.Println(a, b, c)

	_, _, d := fun1()
	fmt.Println(d)

	fmt.Println(fun2(30))
}

func fun1() (int, float64, string) {
	return 0, 0, " test "
}

func fun2(age int) int {
	if age >= 0 {
		return age
	}
	/*{
		fmt.Println("age not right")
		return 0
	}*/
	return 0
}
