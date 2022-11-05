//go:build ignore

package main

import "fmt"

func main() {
	sum := getSum(5)
	fmt.Printf("sum: %v\n", sum)

	sumf := getFibonacci(12)
	fmt.Printf("sumf: %v\n", sumf)
}

func getSum(n int) int {
	if n == 1 {
		return 1
	}
	return getSum(n-1) + n
}

func getFibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return getFibonacci(n-2) + getFibonacci(n-1)
}
