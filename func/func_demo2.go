//go:build ignore

package main

import "fmt"

func main() {
	getSum(10)
	getSum(20)
	getSum(30)

	getAdd(10, 20)
}

func getSum(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("1~%d sum is:%d\n", n, sum)
}

func getAdd(a int, b int) {
	sum := a + b
	fmt.Printf("%d + %d = %d \n", a, b, sum)
}
