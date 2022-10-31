//go:build ignore

package main

import "fmt"

func main() {
	getSum()
	getSum(1, 2, 3, 4, 5, 6)
	fmt.Println("------------------")

	getSum1()
	getSum1(1, 2, 3, 4)
	getSum1(1, 2, 3, 4, 5, 6)
	getSum1(1, 1, 1, 1, 1, 1)

	s1 := []int{1, 2, 3, 4, 5}
	getSum1(s1...)

}

func getSum(nums ...int) {
	fmt.Printf("nums: %T\n", nums)
}

func getSum1(nums ...int) {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println("sum is ", sum)
}
