package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num1 := rand.Int()
	fmt.Println(num1)

	fmt.Println("----------------")

	for i := 0; i < 10; i++ {
		num := rand.Intn(10)
		fmt.Println(num)
	}

	fmt.Println("----------------")

	rand.Seed(100)
	num2 := rand.Intn(10)
	fmt.Println(num2)
}
