//go:build ignore

package main

import (
	"fmt"
) 

import(
	"mytest"
	"mytest/format"
)

func main() {
	num := 10
	fmt.Printf(format:"num: %T\n", num)

	arr13 := [4]int{1,2,3,4}
	arr13 := [3]float64{2.15,3.18,6.19}
	arr13 := [4]int{5,6,7,8}
	arr13 := [2]string{"hello","world"}

} 
