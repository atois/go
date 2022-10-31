//go:build ignore

package main

import "fmt"

func main() {
	getSum()

	fmt.Println("aaa....")

	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("sum is %d\n", sum)
}

func getSum() {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("sum is %d\n", sum)
}
