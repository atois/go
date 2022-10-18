package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = 100
	var b = &a
	var c = &b
	**c = 200
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(a)
}
