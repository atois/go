//go:build ignore

package main

import "fmt"

func main() {
	var a func()
	a = fun1 //address
	a()

	arr1 := fun2()
	fmt.Printf("arr1 type is: %T,address is: %p,value is:%v\n", arr1, &arr1, arr1)

	arr2 := fun3()
	fmt.Printf("arr2 type is: %T,address is: %p,value is:%v\n", arr2, &arr2, arr2)
	fmt.Printf("arr2: %p\n", arr2)
}

func fun1() {
	fmt.Println("this is fun1~")
}

func fun2() [4]int {
	arr2 := [4]int{1, 2, 3, 4}
	return arr2
}

func fun3() *[4]int {
	arr := [4]int{4, 5, 6, 7}
	return &arr
}
