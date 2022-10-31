//go:build ignore

package main

import "fmt"

func main() {
	res := getSum()
	fmt.Println("1-10 = ", res)
	fmt.Println(getSum2())
}

func getSum() int {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	return sum
}

func getSum2() (sum int) {
	fmt.Println("hanshu", sum)
	return 1
}
