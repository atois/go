package main

import "fmt"

func main() {
	var x = 10
	var y = &x
	var z = *y
	x = 20

	fmt.Println(x)
	fmt.Println(*y)
	fmt.Println(z)
}
