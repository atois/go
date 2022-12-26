package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Printf("please type your age")
	fmt.Scan(&age)

	if age > 18 {
		fmt.Printf("you are the man")
	}
}
