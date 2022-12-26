//go:build ignore

package main

import "fmt"

func main() {
	res1, res2 := rectangle(5, 3)
	fmt.Println("zhouchang: ", res1)
	fmt.Println("mianji: ", res2)

}

func rectangle(len, wid float64) (float64, float64) {
	perimeter := (len + wid) * 2
	area := len * wid
	return perimeter, area
}

func getSum2() (sum int) {
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return
}
