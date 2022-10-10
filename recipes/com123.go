package main

import (
	"fmt"
)

func main() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Printf("please type your information\n")
	fmt.Scan(&name, &age, &married)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("married: %v\n", married)
}
